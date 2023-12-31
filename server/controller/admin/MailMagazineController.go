package admin

import (
	"context"
	"errors"

	"server/api/v1/admin"
	adminv1connect "server/api/v1/admin/adminconnect"
	"server/api/v1/shared"
	"server/controller"
	"server/core/entity"
	"server/core/infra/queryService/types"
	usecase "server/core/usecase/admin"

	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MailMagazineController struct {
	magazineUseCase usecase.MailMagazineUsecase
}

var _ adminv1connect.MailMagazineControllerClient = &MailMagazineController{}

func NewMailMagazineController(magazineUsecase *usecase.MailMagazineUsecase) *MailMagazineController {
	return &MailMagazineController{
		magazineUseCase: *magazineUsecase,
	}
}

func (uc *MailMagazineController) GetList(ctx context.Context, req *connect.Request[admin.GetMailMagazineListRequest]) (*connect.Response[admin.MailMagazinesResponse], error) {
	var currentPage, perPage int
	if req.Msg.Pager.CurrentPage != nil {
		currentPage = int(*req.Msg.Pager.CurrentPage)
	}
	if req.Msg.Pager.PerPage != nil {
		perPage = int(*req.Msg.Pager.PerPage)
	}

	pager := &types.PageQuery{
		CurrentPage: &currentPage,
		PerPage:     &perPage,
	}
	entities, domaiErr := uc.magazineUseCase.GetList(pager)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	var mailMagazines []*admin.MailMagazine

	for _, entity := range entities {
		var prefs []shared.Prefecture
		for _, pref := range *entity.TargetPrefecture {
			prefs = append(prefs, shared.Prefecture(pref))
		}

		magazine := &admin.MailMagazine{
			ID:                 entity.ID.String(),
			Title:              entity.Title,
			Content:            entity.Content,
			AuthorID:           entity.AuthorID.String(),
			TargetPrefecture:   prefs,
			MailMagazineStatus: admin.MailMagazineStatus(entity.MailMagazineStatus),
			UnsentCount:        int32(entity.UnsentCount),
			SentCount:          int32(entity.SentCount),
		}
		mailMagazines = append(mailMagazines, magazine)
	}
	result := &admin.MailMagazinesResponse{
		MailMagazines: mailMagazines,
	}
	return connect.NewResponse(result), nil
}

func (uc *MailMagazineController) CreateDraft(ctx context.Context, req *connect.Request[admin.CreateDraftRequest]) (*connect.Response[admin.MailMagazine], error) {
	adminID := ctx.Value("userID").(uuid.UUID)

	if adminID == uuid.Nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	msg := req.Msg
	var prefs []entity.Prefecture
	for _, pref := range msg.TargetPrefectures {
		prefs = append(prefs, entity.Prefecture(pref))
	}
	entity, domainErr := uc.magazineUseCase.CreateDraft(msg.Title, msg.Content, &prefs, adminID)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	var sentAt *timestamppb.Timestamp
	if entity.SentAt != nil {
		sentAt = timestamppb.New(*entity.SentAt)
	}
	return connect.NewResponse(&admin.MailMagazine{
		ID:                 entity.ID.String(),
		Title:              entity.Title,
		Content:            entity.Content,
		AuthorID:           entity.AuthorID.String(),
		TargetPrefecture:   msg.TargetPrefectures,
		MailMagazineStatus: admin.MailMagazineStatus(entity.MailMagazineStatus),
		UnsentCount:        int32(entity.UnsentCount),
		SentCount:          int32(entity.SentCount),
		SentAt:             sentAt,
	}), nil
}

func (uc *MailMagazineController) Update(ctx context.Context, req *connect.Request[admin.UpdateMailMagazineRequest]) (*connect.Response[admin.MailMagazine], error) {
	adminID := ctx.Value("userID").(uuid.UUID)

	if adminID == uuid.Nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	msg := req.Msg
	var entityPrefs []entity.Prefecture
	for _, pref := range msg.TargetPrefectures {
		entityPrefs = append(entityPrefs, entity.Prefecture(pref))
	}
	entity, domaiErr := uc.magazineUseCase.Update(
		msg.Title,
		msg.Content,
		&entityPrefs,
		adminID,
		uuid.MustParse(msg.ID),
	)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	var reqPrefs []shared.Prefecture
	for _, pref := range *entity.TargetPrefecture {
		reqPrefs = append(reqPrefs, shared.Prefecture(pref))
	}

	magazine := &admin.MailMagazine{
		ID:                 entity.ID.String(),
		Title:              entity.Title,
		Content:            entity.Content,
		AuthorID:           entity.AuthorID.String(),
		TargetPrefecture:   reqPrefs,
		MailMagazineStatus: admin.MailMagazineStatus(entity.MailMagazineStatus),
		UnsentCount:        int32(entity.UnsentCount),
		SentCount:          int32(entity.SentCount),
	}

	return connect.NewResponse(magazine), nil
}

func (uc *MailMagazineController) Delete(ctx context.Context, req *connect.Request[admin.DeleteMailMagazineRequest]) (*connect.Response[emptypb.Empty], error) {
	domainErr := uc.magazineUseCase.Delete(uuid.MustParse(req.Msg.ID))
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (uc *MailMagazineController) Send(ctx context.Context, req *connect.Request[admin.SendMailMagazineRequest]) (*connect.Response[emptypb.Empty], error) {
	domainErr := uc.magazineUseCase.Send(uuid.MustParse(req.Msg.ID))
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
