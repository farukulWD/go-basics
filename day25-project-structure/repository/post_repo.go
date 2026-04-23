package repository

import (
	"go-basics/day25-project-structure/domain"

	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) domain.PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *domain.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) FindBySlug(slug string) (*domain.Post, error) {
	var post domain.Post
	err := r.db.Where("slug = ?", slug).First(&post).Error
	return &post, err
}

func (r *postRepository) FindByID(id uint) (*domain.Post, error) {
	var post domain.Post
	err := r.db.First(&post, id).Error
	return &post, err
}

func (r *postRepository) Update(post *domain.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Post{}, id).Error
}

func (r *postRepository) FindPublished(params domain.PaginationParams) ([]domain.Post, int64, error) {
	var posts []domain.Post
	var total int64

	query := r.db.Model(&domain.Post{}).Where("published = true")
	query.Count(&total)

	offset := (params.Page - 1) * params.Limit
	err := query.Order("published_at DESC").
		Limit(params.Limit).
		Offset(offset).
		Find(&posts).Error

	return posts, total, err
}

func (r *postRepository) RecordEvent(event *domain.PostAnalytic) error {
	return r.db.Create(event).Error
}

func (r *postRepository) CountEvents(postID uint, eventType string) (int64, error) {
	var count int64
	err := r.db.Model(&domain.PostAnalytic{}).
		Where("post_id = ? AND event = ?", postID, eventType).
		Count(&count).Error
	return count, err
}

func (r *postRepository) RecentEvents(postID uint, limit int) ([]domain.PostAnalytic, error) {
	var events []domain.PostAnalytic
	err := r.db.Where("post_id = ?", postID).
		Order("created_at DESC").
		Limit(limit).
		Find(&events).Error
	return events, err
}
