package models

import "gorm.io/gorm"


type Post struct {
    // gorm.Model会添加create_at、delete_at等字段，并且提供软删除功能
    gorm.Model
    Title string
    Body string
}
