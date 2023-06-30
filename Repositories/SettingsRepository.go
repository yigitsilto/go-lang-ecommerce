package Repositories

import (
	"ecommerce/dto"
	"gorm.io/gorm"
)

type SettingsRepositoryInterface interface {
	GetSettings() (dto.GeneralSettingsModel, error)
}

type SettingsRepositoryImpl struct {
	db *gorm.DB
}

func NewSettingsRepository(db *gorm.DB) SettingsRepositoryInterface {
	return &SettingsRepositoryImpl{db: db}
}

func (s *SettingsRepositoryImpl) GetSettings() (dto.GeneralSettingsModel, error) {

	settingsKey := [...]string{
		"storefront_copyright_text",
		"storefront_welcome_text",
		"storefront_header_logo",
		"storefront_slider_banner_1_file_id",
		"storefront_slider_banner_2_file_id",
		"storefront_custom_theme_color",
		"storefront_mail_theme_color",
		"storefront_favicon",
		"storefront_facebook_link",
		"storefront_twitter_link",
		"storefront_instagram_link",
		"storefront_youtube_link",
		"storefront_navbar_text",
		"storefront_primary_menu",
		"storefront_category_menu",
		"storefront_footer_menu_one",
		"storefront_footer_menu_one_title",
		"storefront_footer_menu_two",
		"storefront_footer_menu_two_title",
		"popular_products_text",
		"popular_categories_text",
		"related_products_text",
		"blog_text",
	}

	var settingsModel []dto.SettingsModel

	err := s.db.Table("storefront_settings").
		Select("storefront_settings.key, storefront_settings.value").
		Where("storefront_settings.key IN (?)", settingsKey).Find(&settingsModel).Error

	settingsMap := make(map[string]string, len(settingsModel))
	for _, item := range settingsModel {
		settingsMap[item.Key] = item.Value
	}

	menus, err := s.getMenus(settingsMap)
	footer1, err := s.getFooter1(settingsMap)
	footer2, err := s.getFooter2(settingsMap)

	// TODO menü child parent eklenecek
	//menuTree, err := s.buildMenuTrees(menus)

	generalSettings := dto.GeneralSettingsModel{Settings: settingsMap, Menu: menus, Footer1: footer1, Footer2: footer2}

	return generalSettings, err

}

func (s *SettingsRepositoryImpl) getMenus(settingsMap map[string]string) ([]dto.MenuModel, error) {
	var menus []dto.MenuModel

	err := s.db.Table("menus").
		Select("mt.name, mi.parent_id, mi.is_root, mi.url, mi.type, mi.parent_id, mi.id  ").
		Joins(
			"INNER JOIN menu_items mi ON mi.menu_id = menus.id AND mi.is_active =? "+
				"INNER JOIN menu_item_translations mt ON mt.menu_item_id = mi.id ", true,
		).
		Where(
			"menus.is_active =? AND mt.name != ? AND menus.id =? AND mi.parent_id = ?", true, "root",
			settingsMap["storefront_primary_menu"], 6,
		).
		Find(&menus).
		Error

	return menus, err

}

func (s *SettingsRepositoryImpl) getFooter1(settingsMap map[string]string) ([]dto.MenuModel, error) {
	var menus []dto.MenuModel

	err := s.db.Table("menus").
		Select("mt.name, mi.parent_id, mi.is_root, mi.url, mi.type, mi.parent_id, mi.id  ").
		Joins(
			"INNER JOIN menu_items mi ON mi.menu_id = menus.id AND mi.is_active =? "+
				"INNER JOIN menu_item_translations mt ON mt.menu_item_id = mi.id ", true,
		).
		Where(
			"menus.is_active =? AND mt.name != ? AND menus.id =? AND mi.parent_id = ?", true, "root",
			settingsMap["storefront_footer_menu_one"], 9,
		).
		Find(&menus).
		Error

	return menus, err

}

func (s *SettingsRepositoryImpl) getFooter2(settingsMap map[string]string) ([]dto.MenuModel, error) {
	var menus []dto.MenuModel

	err := s.db.Table("menus").
		Select("mt.name, mi.parent_id, mi.is_root, mi.url, mi.type, mi.parent_id, mi.id  ").
		Joins(
			"INNER JOIN menu_items mi ON mi.menu_id = menus.id AND mi.is_active =? "+
				"INNER JOIN menu_item_translations mt ON mt.menu_item_id = mi.id ", true,
		).
		Where(
			"menus.is_active =? AND mt.name != ? AND menus.id =? AND mi.parent_id = ?", true, "root",
			settingsMap["storefront_footer_menu_two"], 12,
		).
		Find(&menus).
		Error

	return menus, err

}

/*
func (s *SettingsRepositoryImpl) buildMenuTrees(menus []dto.MenuModel) []dto.MenuModel {

	var child dto.SubMenuModel

	for _, item := range menus {

		if item.IsRoot {

		}

	}

	return nil
}
*/
