// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Admins", testAdmins)
	t.Run("Banners", testBanners)
	t.Run("Checkins", testCheckins)
	t.Run("Coupons", testCoupons)
	t.Run("CouponAttachedUsers", testCouponAttachedUsers)
	t.Run("CouponNotices", testCouponNotices)
	t.Run("CouponStores", testCouponStores)
	t.Run("MailMagazines", testMailMagazines)
	t.Run("Posts", testPosts)
	t.Run("SchemaMigrations", testSchemaMigrations)
	t.Run("StayableStoreInfos", testStayableStoreInfos)
	t.Run("Stores", testStores)
	t.Run("Users", testUsers)
	t.Run("UserOptions", testUserOptions)
}

func TestDelete(t *testing.T) {
	t.Run("Admins", testAdminsDelete)
	t.Run("Banners", testBannersDelete)
	t.Run("Checkins", testCheckinsDelete)
	t.Run("Coupons", testCouponsDelete)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersDelete)
	t.Run("CouponNotices", testCouponNoticesDelete)
	t.Run("CouponStores", testCouponStoresDelete)
	t.Run("MailMagazines", testMailMagazinesDelete)
	t.Run("Posts", testPostsDelete)
	t.Run("SchemaMigrations", testSchemaMigrationsDelete)
	t.Run("StayableStoreInfos", testStayableStoreInfosDelete)
	t.Run("Stores", testStoresDelete)
	t.Run("Users", testUsersDelete)
	t.Run("UserOptions", testUserOptionsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Admins", testAdminsQueryDeleteAll)
	t.Run("Banners", testBannersQueryDeleteAll)
	t.Run("Checkins", testCheckinsQueryDeleteAll)
	t.Run("Coupons", testCouponsQueryDeleteAll)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersQueryDeleteAll)
	t.Run("CouponNotices", testCouponNoticesQueryDeleteAll)
	t.Run("CouponStores", testCouponStoresQueryDeleteAll)
	t.Run("MailMagazines", testMailMagazinesQueryDeleteAll)
	t.Run("Posts", testPostsQueryDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsQueryDeleteAll)
	t.Run("StayableStoreInfos", testStayableStoreInfosQueryDeleteAll)
	t.Run("Stores", testStoresQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
	t.Run("UserOptions", testUserOptionsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Admins", testAdminsSliceDeleteAll)
	t.Run("Banners", testBannersSliceDeleteAll)
	t.Run("Checkins", testCheckinsSliceDeleteAll)
	t.Run("Coupons", testCouponsSliceDeleteAll)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersSliceDeleteAll)
	t.Run("CouponNotices", testCouponNoticesSliceDeleteAll)
	t.Run("CouponStores", testCouponStoresSliceDeleteAll)
	t.Run("MailMagazines", testMailMagazinesSliceDeleteAll)
	t.Run("Posts", testPostsSliceDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceDeleteAll)
	t.Run("StayableStoreInfos", testStayableStoreInfosSliceDeleteAll)
	t.Run("Stores", testStoresSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
	t.Run("UserOptions", testUserOptionsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Admins", testAdminsExists)
	t.Run("Banners", testBannersExists)
	t.Run("Checkins", testCheckinsExists)
	t.Run("Coupons", testCouponsExists)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersExists)
	t.Run("CouponNotices", testCouponNoticesExists)
	t.Run("CouponStores", testCouponStoresExists)
	t.Run("MailMagazines", testMailMagazinesExists)
	t.Run("Posts", testPostsExists)
	t.Run("SchemaMigrations", testSchemaMigrationsExists)
	t.Run("StayableStoreInfos", testStayableStoreInfosExists)
	t.Run("Stores", testStoresExists)
	t.Run("Users", testUsersExists)
	t.Run("UserOptions", testUserOptionsExists)
}

func TestFind(t *testing.T) {
	t.Run("Admins", testAdminsFind)
	t.Run("Banners", testBannersFind)
	t.Run("Checkins", testCheckinsFind)
	t.Run("Coupons", testCouponsFind)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersFind)
	t.Run("CouponNotices", testCouponNoticesFind)
	t.Run("CouponStores", testCouponStoresFind)
	t.Run("MailMagazines", testMailMagazinesFind)
	t.Run("Posts", testPostsFind)
	t.Run("SchemaMigrations", testSchemaMigrationsFind)
	t.Run("StayableStoreInfos", testStayableStoreInfosFind)
	t.Run("Stores", testStoresFind)
	t.Run("Users", testUsersFind)
	t.Run("UserOptions", testUserOptionsFind)
}

func TestBind(t *testing.T) {
	t.Run("Admins", testAdminsBind)
	t.Run("Banners", testBannersBind)
	t.Run("Checkins", testCheckinsBind)
	t.Run("Coupons", testCouponsBind)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersBind)
	t.Run("CouponNotices", testCouponNoticesBind)
	t.Run("CouponStores", testCouponStoresBind)
	t.Run("MailMagazines", testMailMagazinesBind)
	t.Run("Posts", testPostsBind)
	t.Run("SchemaMigrations", testSchemaMigrationsBind)
	t.Run("StayableStoreInfos", testStayableStoreInfosBind)
	t.Run("Stores", testStoresBind)
	t.Run("Users", testUsersBind)
	t.Run("UserOptions", testUserOptionsBind)
}

func TestOne(t *testing.T) {
	t.Run("Admins", testAdminsOne)
	t.Run("Banners", testBannersOne)
	t.Run("Checkins", testCheckinsOne)
	t.Run("Coupons", testCouponsOne)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersOne)
	t.Run("CouponNotices", testCouponNoticesOne)
	t.Run("CouponStores", testCouponStoresOne)
	t.Run("MailMagazines", testMailMagazinesOne)
	t.Run("Posts", testPostsOne)
	t.Run("SchemaMigrations", testSchemaMigrationsOne)
	t.Run("StayableStoreInfos", testStayableStoreInfosOne)
	t.Run("Stores", testStoresOne)
	t.Run("Users", testUsersOne)
	t.Run("UserOptions", testUserOptionsOne)
}

func TestAll(t *testing.T) {
	t.Run("Admins", testAdminsAll)
	t.Run("Banners", testBannersAll)
	t.Run("Checkins", testCheckinsAll)
	t.Run("Coupons", testCouponsAll)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersAll)
	t.Run("CouponNotices", testCouponNoticesAll)
	t.Run("CouponStores", testCouponStoresAll)
	t.Run("MailMagazines", testMailMagazinesAll)
	t.Run("Posts", testPostsAll)
	t.Run("SchemaMigrations", testSchemaMigrationsAll)
	t.Run("StayableStoreInfos", testStayableStoreInfosAll)
	t.Run("Stores", testStoresAll)
	t.Run("Users", testUsersAll)
	t.Run("UserOptions", testUserOptionsAll)
}

func TestCount(t *testing.T) {
	t.Run("Admins", testAdminsCount)
	t.Run("Banners", testBannersCount)
	t.Run("Checkins", testCheckinsCount)
	t.Run("Coupons", testCouponsCount)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersCount)
	t.Run("CouponNotices", testCouponNoticesCount)
	t.Run("CouponStores", testCouponStoresCount)
	t.Run("MailMagazines", testMailMagazinesCount)
	t.Run("Posts", testPostsCount)
	t.Run("SchemaMigrations", testSchemaMigrationsCount)
	t.Run("StayableStoreInfos", testStayableStoreInfosCount)
	t.Run("Stores", testStoresCount)
	t.Run("Users", testUsersCount)
	t.Run("UserOptions", testUserOptionsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Admins", testAdminsHooks)
	t.Run("Banners", testBannersHooks)
	t.Run("Checkins", testCheckinsHooks)
	t.Run("Coupons", testCouponsHooks)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersHooks)
	t.Run("CouponNotices", testCouponNoticesHooks)
	t.Run("CouponStores", testCouponStoresHooks)
	t.Run("MailMagazines", testMailMagazinesHooks)
	t.Run("Posts", testPostsHooks)
	t.Run("SchemaMigrations", testSchemaMigrationsHooks)
	t.Run("StayableStoreInfos", testStayableStoreInfosHooks)
	t.Run("Stores", testStoresHooks)
	t.Run("Users", testUsersHooks)
	t.Run("UserOptions", testUserOptionsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Admins", testAdminsInsert)
	t.Run("Admins", testAdminsInsertWhitelist)
	t.Run("Banners", testBannersInsert)
	t.Run("Banners", testBannersInsertWhitelist)
	t.Run("Checkins", testCheckinsInsert)
	t.Run("Checkins", testCheckinsInsertWhitelist)
	t.Run("Coupons", testCouponsInsert)
	t.Run("Coupons", testCouponsInsertWhitelist)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersInsert)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersInsertWhitelist)
	t.Run("CouponNotices", testCouponNoticesInsert)
	t.Run("CouponNotices", testCouponNoticesInsertWhitelist)
	t.Run("CouponStores", testCouponStoresInsert)
	t.Run("CouponStores", testCouponStoresInsertWhitelist)
	t.Run("MailMagazines", testMailMagazinesInsert)
	t.Run("MailMagazines", testMailMagazinesInsertWhitelist)
	t.Run("Posts", testPostsInsert)
	t.Run("Posts", testPostsInsertWhitelist)
	t.Run("SchemaMigrations", testSchemaMigrationsInsert)
	t.Run("SchemaMigrations", testSchemaMigrationsInsertWhitelist)
	t.Run("StayableStoreInfos", testStayableStoreInfosInsert)
	t.Run("StayableStoreInfos", testStayableStoreInfosInsertWhitelist)
	t.Run("Stores", testStoresInsert)
	t.Run("Stores", testStoresInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
	t.Run("UserOptions", testUserOptionsInsert)
	t.Run("UserOptions", testUserOptionsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("AdminToStoreUsingBelongToStore", testAdminToOneStoreUsingBelongToStore)
	t.Run("CheckinToStoreUsingStore", testCheckinToOneStoreUsingStore)
	t.Run("CheckinToUserUsingUser", testCheckinToOneUserUsingUser)
	t.Run("CouponAttachedUserToCouponUsingCoupon", testCouponAttachedUserToOneCouponUsingCoupon)
	t.Run("CouponAttachedUserToUserUsingUser", testCouponAttachedUserToOneUserUsingUser)
	t.Run("CouponNoticeToCouponUsingCoupon", testCouponNoticeToOneCouponUsingCoupon)
	t.Run("CouponStoreToCouponUsingCoupon", testCouponStoreToOneCouponUsingCoupon)
	t.Run("CouponStoreToStoreUsingStore", testCouponStoreToOneStoreUsingStore)
	t.Run("MailMagazineToAdminUsingAuthorAdmin", testMailMagazineToOneAdminUsingAuthorAdmin)
	t.Run("PostToAdminUsingAuthorAdmin", testPostToOneAdminUsingAuthorAdmin)
	t.Run("StayableStoreInfoToStoreUsingStore", testStayableStoreInfoToOneStoreUsingStore)
	t.Run("UserOptionToUserUsingUser", testUserOptionToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {
	t.Run("CouponToCouponNoticeUsingCouponNotice", testCouponOneToOneCouponNoticeUsingCouponNotice)
	t.Run("CouponToCouponStoreUsingCouponStore", testCouponOneToOneCouponStoreUsingCouponStore)
	t.Run("StoreToStayableStoreInfoUsingStayableStoreInfo", testStoreOneToOneStayableStoreInfoUsingStayableStoreInfo)
}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AdminToAuthorMailMagazines", testAdminToManyAuthorMailMagazines)
	t.Run("AdminToAuthorPosts", testAdminToManyAuthorPosts)
	t.Run("CouponToCouponAttachedUsers", testCouponToManyCouponAttachedUsers)
	t.Run("StoreToBelongToAdmins", testStoreToManyBelongToAdmins)
	t.Run("StoreToCheckins", testStoreToManyCheckins)
	t.Run("StoreToCouponStores", testStoreToManyCouponStores)
	t.Run("UserToCheckins", testUserToManyCheckins)
	t.Run("UserToCouponAttachedUsers", testUserToManyCouponAttachedUsers)
	t.Run("UserToUserOptions", testUserToManyUserOptions)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("AdminToStoreUsingBelongToAdmins", testAdminToOneSetOpStoreUsingBelongToStore)
	t.Run("CheckinToStoreUsingCheckins", testCheckinToOneSetOpStoreUsingStore)
	t.Run("CheckinToUserUsingCheckins", testCheckinToOneSetOpUserUsingUser)
	t.Run("CouponAttachedUserToCouponUsingCouponAttachedUsers", testCouponAttachedUserToOneSetOpCouponUsingCoupon)
	t.Run("CouponAttachedUserToUserUsingCouponAttachedUsers", testCouponAttachedUserToOneSetOpUserUsingUser)
	t.Run("CouponNoticeToCouponUsingCouponNotice", testCouponNoticeToOneSetOpCouponUsingCoupon)
	t.Run("CouponStoreToCouponUsingCouponStore", testCouponStoreToOneSetOpCouponUsingCoupon)
	t.Run("CouponStoreToStoreUsingCouponStores", testCouponStoreToOneSetOpStoreUsingStore)
	t.Run("MailMagazineToAdminUsingAuthorMailMagazines", testMailMagazineToOneSetOpAdminUsingAuthorAdmin)
	t.Run("PostToAdminUsingAuthorPosts", testPostToOneSetOpAdminUsingAuthorAdmin)
	t.Run("StayableStoreInfoToStoreUsingStayableStoreInfo", testStayableStoreInfoToOneSetOpStoreUsingStore)
	t.Run("UserOptionToUserUsingUserOptions", testUserOptionToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("AdminToStoreUsingBelongToAdmins", testAdminToOneRemoveOpStoreUsingBelongToStore)
	t.Run("CheckinToStoreUsingCheckins", testCheckinToOneRemoveOpStoreUsingStore)
	t.Run("CheckinToUserUsingCheckins", testCheckinToOneRemoveOpUserUsingUser)
	t.Run("MailMagazineToAdminUsingAuthorMailMagazines", testMailMagazineToOneRemoveOpAdminUsingAuthorAdmin)
	t.Run("PostToAdminUsingAuthorPosts", testPostToOneRemoveOpAdminUsingAuthorAdmin)
	t.Run("UserOptionToUserUsingUserOptions", testUserOptionToOneRemoveOpUserUsingUser)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {
	t.Run("CouponToCouponNoticeUsingCouponNotice", testCouponOneToOneSetOpCouponNoticeUsingCouponNotice)
	t.Run("CouponToCouponStoreUsingCouponStore", testCouponOneToOneSetOpCouponStoreUsingCouponStore)
	t.Run("StoreToStayableStoreInfoUsingStayableStoreInfo", testStoreOneToOneSetOpStayableStoreInfoUsingStayableStoreInfo)
}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AdminToAuthorMailMagazines", testAdminToManyAddOpAuthorMailMagazines)
	t.Run("AdminToAuthorPosts", testAdminToManyAddOpAuthorPosts)
	t.Run("CouponToCouponAttachedUsers", testCouponToManyAddOpCouponAttachedUsers)
	t.Run("StoreToBelongToAdmins", testStoreToManyAddOpBelongToAdmins)
	t.Run("StoreToCheckins", testStoreToManyAddOpCheckins)
	t.Run("StoreToCouponStores", testStoreToManyAddOpCouponStores)
	t.Run("UserToCheckins", testUserToManyAddOpCheckins)
	t.Run("UserToCouponAttachedUsers", testUserToManyAddOpCouponAttachedUsers)
	t.Run("UserToUserOptions", testUserToManyAddOpUserOptions)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("AdminToAuthorMailMagazines", testAdminToManySetOpAuthorMailMagazines)
	t.Run("AdminToAuthorPosts", testAdminToManySetOpAuthorPosts)
	t.Run("StoreToBelongToAdmins", testStoreToManySetOpBelongToAdmins)
	t.Run("StoreToCheckins", testStoreToManySetOpCheckins)
	t.Run("UserToCheckins", testUserToManySetOpCheckins)
	t.Run("UserToUserOptions", testUserToManySetOpUserOptions)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("AdminToAuthorMailMagazines", testAdminToManyRemoveOpAuthorMailMagazines)
	t.Run("AdminToAuthorPosts", testAdminToManyRemoveOpAuthorPosts)
	t.Run("StoreToBelongToAdmins", testStoreToManyRemoveOpBelongToAdmins)
	t.Run("StoreToCheckins", testStoreToManyRemoveOpCheckins)
	t.Run("UserToCheckins", testUserToManyRemoveOpCheckins)
	t.Run("UserToUserOptions", testUserToManyRemoveOpUserOptions)
}

func TestReload(t *testing.T) {
	t.Run("Admins", testAdminsReload)
	t.Run("Banners", testBannersReload)
	t.Run("Checkins", testCheckinsReload)
	t.Run("Coupons", testCouponsReload)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersReload)
	t.Run("CouponNotices", testCouponNoticesReload)
	t.Run("CouponStores", testCouponStoresReload)
	t.Run("MailMagazines", testMailMagazinesReload)
	t.Run("Posts", testPostsReload)
	t.Run("SchemaMigrations", testSchemaMigrationsReload)
	t.Run("StayableStoreInfos", testStayableStoreInfosReload)
	t.Run("Stores", testStoresReload)
	t.Run("Users", testUsersReload)
	t.Run("UserOptions", testUserOptionsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Admins", testAdminsReloadAll)
	t.Run("Banners", testBannersReloadAll)
	t.Run("Checkins", testCheckinsReloadAll)
	t.Run("Coupons", testCouponsReloadAll)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersReloadAll)
	t.Run("CouponNotices", testCouponNoticesReloadAll)
	t.Run("CouponStores", testCouponStoresReloadAll)
	t.Run("MailMagazines", testMailMagazinesReloadAll)
	t.Run("Posts", testPostsReloadAll)
	t.Run("SchemaMigrations", testSchemaMigrationsReloadAll)
	t.Run("StayableStoreInfos", testStayableStoreInfosReloadAll)
	t.Run("Stores", testStoresReloadAll)
	t.Run("Users", testUsersReloadAll)
	t.Run("UserOptions", testUserOptionsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Admins", testAdminsSelect)
	t.Run("Banners", testBannersSelect)
	t.Run("Checkins", testCheckinsSelect)
	t.Run("Coupons", testCouponsSelect)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersSelect)
	t.Run("CouponNotices", testCouponNoticesSelect)
	t.Run("CouponStores", testCouponStoresSelect)
	t.Run("MailMagazines", testMailMagazinesSelect)
	t.Run("Posts", testPostsSelect)
	t.Run("SchemaMigrations", testSchemaMigrationsSelect)
	t.Run("StayableStoreInfos", testStayableStoreInfosSelect)
	t.Run("Stores", testStoresSelect)
	t.Run("Users", testUsersSelect)
	t.Run("UserOptions", testUserOptionsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Admins", testAdminsUpdate)
	t.Run("Banners", testBannersUpdate)
	t.Run("Checkins", testCheckinsUpdate)
	t.Run("Coupons", testCouponsUpdate)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersUpdate)
	t.Run("CouponNotices", testCouponNoticesUpdate)
	t.Run("CouponStores", testCouponStoresUpdate)
	t.Run("MailMagazines", testMailMagazinesUpdate)
	t.Run("Posts", testPostsUpdate)
	t.Run("SchemaMigrations", testSchemaMigrationsUpdate)
	t.Run("StayableStoreInfos", testStayableStoreInfosUpdate)
	t.Run("Stores", testStoresUpdate)
	t.Run("Users", testUsersUpdate)
	t.Run("UserOptions", testUserOptionsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Admins", testAdminsSliceUpdateAll)
	t.Run("Banners", testBannersSliceUpdateAll)
	t.Run("Checkins", testCheckinsSliceUpdateAll)
	t.Run("Coupons", testCouponsSliceUpdateAll)
	t.Run("CouponAttachedUsers", testCouponAttachedUsersSliceUpdateAll)
	t.Run("CouponNotices", testCouponNoticesSliceUpdateAll)
	t.Run("CouponStores", testCouponStoresSliceUpdateAll)
	t.Run("MailMagazines", testMailMagazinesSliceUpdateAll)
	t.Run("Posts", testPostsSliceUpdateAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceUpdateAll)
	t.Run("StayableStoreInfos", testStayableStoreInfosSliceUpdateAll)
	t.Run("Stores", testStoresSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
	t.Run("UserOptions", testUserOptionsSliceUpdateAll)
}
