package domain

import "time"

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string     `gorm:"not null"`
	Slug        string     `gorm:"uniqueIndex;not null"`
	Content     string     `gorm:"type:text;not null"`
	AuthorID    uint       `gorm:"not null"`
	Published   bool       `gorm:"default:false"`
	PublishedAt *time.Time
}

type PostAnalytic struct {
	gorm.Model
	PostID    uint   `gorm:"not null;index"`
	Event     string `gorm:"not null"`
	UserID    *uint
	IPAddress string
	UserAgent string
}

type PostResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Content     string     `json:"content"`
	AuthorID    uint       `json:"author_id"`
	Published   bool       `json:"published"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
}

func (p *Post) ToResponse() PostResponse {
	return PostResponse{
		ID:          p.ID,
		Title:       p.Title,
		Slug:        p.Slug,
		Content:     p.Content,
		AuthorID:    p.AuthorID,
		Published:   p.Published,
		PublishedAt: p.PublishedAt,
		CreatedAt:   p.CreatedAt,
	}
}

type AnalyticEvent struct {
	ID        uint      `json:"id"`
	Event     string    `json:"event"`
	UserID    *uint     `json:"user_id"`
	IPAddress string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	CreatedAt time.Time `json:"created_at"`
}

type AnalyticsSummary struct {
	PostID       uint            `json:"post_id"`
	ViewCount    int64           `json:"view_count"`
	LikeCount    int64           `json:"like_count"`
	RecentEvents []AnalyticEvent `json:"recent_events"`
}

type PaginationParams struct {
	Page  int
	Limit int
}

type PaginatedPosts struct {
	Data       []PostResponse `json:"data"`
	Total      int64          `json:"total"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalPages int            `json:"total_pages"`
}

type PostRepository interface {
	Create(post *Post) error
	FindBySlug(slug string) (*Post, error)
	FindByID(id uint) (*Post, error)
	Update(post *Post) error
	Delete(id uint) error
	FindPublished(params PaginationParams) ([]Post, int64, error)
	RecordEvent(event *PostAnalytic) error
	CountEvents(postID uint, eventType string) (int64, error)
	RecentEvents(postID uint, limit int) ([]PostAnalytic, error)
}

type PostService interface {
	CreatePost(authorID uint, title, content string) (*PostResponse, error)
	UpdatePost(id uint, title, content string, published bool) (*PostResponse, error)
	DeletePost(id uint) error
	GetBySlug(slug, ip, userAgent string, userID *uint) (*PostResponse, error)
	ListPublished(page, limit int) (*PaginatedPosts, error)
	LikePost(slug, ip, userAgent string, userID *uint) error
	GetAnalytics(postID uint) (*AnalyticsSummary, error)
}
