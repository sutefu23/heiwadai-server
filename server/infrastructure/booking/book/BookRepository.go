package book

import (
	"errors"
	"time"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
	"server/infrastructure/booking/util"
	"server/infrastructure/env"
	"server/infrastructure/logger"

	"github.com/ktnyt/go-moji"
)

type BookRepository struct {
	storeQuery queryservice.IStoreQueryService
	bookQuery  queryservice.IBookQueryService
}

var _ repository.IBookAPIRepository = &BookRepository{}
var (
	TLBookingUser = env.GetEnv(env.TlbookingUsername)
	TLBookingPass = env.GetEnv(env.TlbookingPassword)
	BookURL       = env.GetEnv(env.TlbookingBookingApiUrl)
	CancelURL     = env.GetEnv(env.TlbookingCancelApiUrl)
)

func NewBookRepository(storeQuery queryservice.IStoreQueryService, bookQuery queryservice.IBookQueryService) *BookRepository {
	return &BookRepository{
		storeQuery: storeQuery,
		bookQuery:  bookQuery,
	}
}

func (p BookRepository) Cancel(bookData *entity.Booking) error {
	store, err := p.storeQuery.GetStayableByID(bookData.BookPlan.StoreID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	dataID, err := p.bookQuery.GetBookRequestDataID()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	reqBody := NewCancelRQ(bookData, store, *dataID)
	res, err := util.Request[EnvelopeRQ[CancelBody], EnvelopeRS[CancelBodyRS]](CancelURL, reqBody)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if res.Body.DeleteBookingWithCPResponse.DeleteBookingWithCPResult.CommonResponse.ResultCode == "False" {
		msg := res.Body.DeleteBookingWithCPResponse.DeleteBookingWithCPResult.CommonResponse.ErrorInfos.ErrorMsg
		code := res.Body.DeleteBookingWithCPResponse.DeleteBookingWithCPResult.CommonResponse.ErrorInfos.ErrorCode
		logger.Error(code + ":" + msg)
		return errors.New(msg)
	}

	return nil
}

func (p *BookRepository) Reserve(
	bookData *entity.Booking,
) (*string, error) {
	store, err := p.storeQuery.GetStayableByID(bookData.BookPlan.StoreID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	dataID, err := p.bookQuery.GetBookRequestDataID()
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	reqBody := NewBookingRQ(
		bookData, store, *dataID,
	)

	res, err := util.Request[EnvelopeRQ[Body], EnvelopeRS[BodyRS]](BookURL, reqBody)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if res.Body.EntryBookingResponse.EntryBookingResult.CommonResponse.ResultCode == "False" {
		msg := res.Body.EntryBookingResponse.EntryBookingResult.CommonResponse.ErrorInfos.ErrorMsg
		code := res.Body.EntryBookingResponse.EntryBookingResult.CommonResponse.ErrorInfos.ErrorCode
		logger.Error(code + ":" + msg)
		return nil, errors.New(msg)
	}

	TlBookingNumber := res.Body.EntryBookingResponse.EntryBookingResult.ExtendLincoln.TllBookingNumber
	return TlBookingNumber, nil
}

// bookID : CCYYMMDD+9桁連番（0埋め、データ毎に+1）
func NewBookingRQ(bookData *entity.Booking, store *entity.StayableStore, dataID string) *EnvelopeRQ[Body] {
	plan := bookData.BookPlan
	guest := bookData.GuestData

	var mealCondition MealCondition

	switch plan.MealType.String() {
	case "朝食あり夕食あり":
		mealCondition = MealCondition1night2meals
	case "朝食あり夕食なし":
		mealCondition = MealCondition1nightBreakfast
	case "朝食なし夕食あり":
		mealCondition = MealCondition1nightBreakfast
	case "食事なし":
		mealCondition = MealConditionWithoutMeal
	default:
		mealCondition = MealConditionOther
	}

	var BookingSystemID string
	if env.GetEnv(env.TlbookingIsTest) == "true" {
		BookingSystemID = "E69502"
	} else {
		BookingSystemID = store.BookingSystemID
	}

	guestNameKana := moji.Convert(guest.LastNameKana, moji.ZK, moji.HK) + " " + moji.Convert(guest.FirstNameKana, moji.ZK, moji.HK)
	return &EnvelopeRQ[Body]{
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Naif:    "http://naifc3000.naifc30.nai.lincoln.seanuts.co.jp/",
		Header:  "",
		Body: Body{
			EntryBooking: EntryBooking{
				EntryBookingRequest: EntryBookingRequest{
					CommonRequest: CommonRequest{
						AgtID:       TLBookingUser,
						AgtPassword: TLBookingPass,
						SystemDate:  time.Now().Format("2006-01-02T15:04:05"),
					},
					ExtendLincoln: ExtendLincoln{
						TllHotelCode: BookingSystemID,
						UseTllPlan:   1,
					},
					SendInformation: SendInformation{
						AssignDiv: 1, // 部屋割ありデフォルト
						GenderDiv: 0, // 男女区分なしデフォルト
					},
					AllotmentBookingReport: AllotmentBookingReport{
						TransactionType: TransactionType{
							DataFrom:           "FromTravelAgency",
							DataClassification: "NewBookReport",
							DataID:             dataID,
						},
						AccommodationInformation: AccommodationInformation{ // 宿泊施設情報
							AccommodationName: store.Name + *store.BranchName,
							AccommodationCode: store.BookingSystemID,
						},
						SalesOfficeInformation: SalesOfficeInformation{
							SalesOfficeCompanyName: "平和台ホテルアプリ",
							SalesOfficeName:        store.Name + *store.BranchName,
							SalesOfficeCode:        "10000000",
						},
						BasicInformation: BasicInformation{
							TravelAgencyBookingNumber:  bookData.TlBookingNumber,
							TravelAgencyBookingDate:    time.Now().Format("2006-01-02"),
							TravelAgencyBookingTime:    time.Now().Format("15:04:05"),
							GuestOrGroupNameSingleByte: guestNameKana,
							GuestOrGroupKanjiName:      guest.LastName + " " + guest.FirstName,
							GuestOrGroupPhoneNumber:    *guest.Tel,
							GuestOrGroupEmail:          guest.Mail,
							GuestOrGroupPostalCode:     *guest.ZipCode,
							GuestOrGroupAddress:        guest.Prefecture.String() + *guest.City + *guest.Address,
							CheckInDate:                util.YYYYMMDD(bookData.StayFrom.Format("2006-01-02")),
							CheckInTime:                string(bookData.CheckInTime),
							Nights:                     uint(bookData.StayTo.Sub(bookData.StayFrom).Hours() / 24),
							TotalRoomCount:             bookData.RoomCount,
							GrandTotalPaxCount:         bookData.Adult + bookData.Child,
							TotalPaxMaleCount:          bookData.Adult,
							TotalPaxFemaleCount:        0,
							TotalChildA70Count:         bookData.Child,
							MealCondition:              mealCondition,
							PackageType:                "Package",
							PackagePlanCode:            plan.ID,
							PackagePlanName:            plan.Title,
						},
						BasicRateInformation: BasicRateInformation{
							RoomRateOrPersonalRate:   RoomRateRoom,
							TaxServiceFee:            IncludingServiceAndTax,
							Payment:                  "Cash",
							SettlementDiv:            0, // 現地決済
							TotalAccommodationCharge: int(bookData.TotalCost),
							PointsDiscountList: PointsDiscountList{
								PointsDiscount: 0, // ポイント値引き
							},
						},
						RoomInformationList: RoomInformationList{
							RoomAndGuestList: []RoomAndGuest{
								{
									RoomInformation: RoomInformation{
										RoomTypeCode:    int(plan.RoomType),
										RoomTypeName:    plan.RoomType.String(),
										PerRoomPaxCount: bookData.Adult + bookData.Child,
									},
									RoomRateInformation: RoomRateInformation{
										RoomDate: util.YYYYMMDD(bookData.StayFrom.Format("2006-01-02")),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func NewCancelRQ(bookData *entity.Booking, store *entity.StayableStore, dataID string) *EnvelopeRQ[CancelBody] {
	return &EnvelopeRQ[CancelBody]{
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Naif:    "http://naifc3000.naifc30.nai.lincoln.seanuts.co.jp/",
		Header:  "",
		Body: CancelBody{
			DeleteBookingWithCP: DeleteBookingWithCP{
				DeleteBookingWithCPRequest: DeleteBookingWithCPRequest{
					CommonRequest: CommonRequest{
						AgtID:       TLBookingUser,
						AgtPassword: TLBookingPass,
						SystemDate:  time.Now().Format("2006-01-02T15:04:05"),
					},
					BookingInfo: BookingInfo{
						TllHotelCode:     store.BookingSystemID,
						TllBookingNumber: bookData.TlBookingNumber,
						DataID:           dataID,
					},
				},
			},
		},
	}
}
