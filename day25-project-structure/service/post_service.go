package service

import (
	"errors"
	"math"
	"regexp"
	"strings"
	"time"

	"go-basics/day25-project-structure/domain"
)

var (
	whitespaceRe = regexp.MustCompile(`\s+`)
	nonSlugRe    = regexp.MustCompile(`[^a-z0-9-]`)
)

func generateSlug(title string) string {
	s := strings.ToLower(title)
	s = whitespaceRe.ReplaceAllString(s, "-")
	s = nonSlugRe.ReplaceAllString(s, "")
	return strings.Trim(s, "-")
}

type postService struct {
	repo domain.PostRepository
}

func NewPostService(repo domain.PostRepository) domain.PostService {
	return &postService{repo: repo}
}

func (s *postService) CreatePost(authorID uint, title, content string) (*domain.PostResponse, error) {
	slug := generateSlug(title)
	post := &domain.Post{
		Title:    title,
		Slug:     slug,
		Content:  content,
		AuthorID: authorID,
	}
	if err := s.repo.Create(post); err != nil {
		return nil, errors.New("could not create post (slug may already exist)")
	}
	res := post.ToResponse()
	return &res, nil
}

func (s *postService) UpdatePost(id uint, title, content string, published bool) (*domain.PostResponse, error) {
	post, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("post not found")
	}
	post.Title = title
	post.Content = content
	if published && !post.Published {
		now := time.Now()
		post.PublishedAt = &now
	}
	post.Published = published
	if err := s.repo.Update(post); err != nil {
		return nil, errors.New("could not update post")
	}
	res := post.ToResponse()
	return &res, nil
}

func (s *postService) DeletePost(id uint) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return errors.New("post not found")
	}
	return s.repo.Delete(id)
}

func (s *postService) GetBySlug(slug, ip, userAgent string, userID *uint) (*domain.PostResponse, error) {
	post, err := s.repo.FindBySlug(slug)
	if err != nil {
		return nil, errors.New("post not found")
	}
	_ = s.repo.RecordEvent(&domain.PostAnalytic{
		PostID:    post.ID,
		Event:     "view",
		UserID:    userID,
		IPAddress: ip,
		UserAgent: userAgent,
	})
	res := post.ToResponse()
	return &res, nil
}

func (s *postService) ListPublished(page, limit int) (*domain.PaginatedPosts, error) {
	posts, total, err := s.repo.FindPublished(domain.PaginationParams{Page: page, Limit: limit})
	if err != nil {
		return nil, err
	}
	var data []domain.PostResponse
	for _, p := range posts {
		data = append(data, p.ToResponse())
	}
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return &domain.PaginatedPosts{
		Data:       data,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

func (s *postService) LikePost(slug, ip, userAgent string, userID *uint) error {
	post, err := s.repo.FindBySlug(slug)
	if err != nil {
		return errors.New("post not found")
	}
	return s.repo.RecordEvent(&domain.PostAnalytic{
		PostID:    post.ID,
		Event:     "like",
		UserID:    userID,
		IPAddress: ip,
		UserAgent: userAgent,
	})
}

func (s *postService) GetAnalytics(postID uint) (*domain.AnalyticsSummary, error) {
	if _, err := s.repo.FindByID(postID); err != nil {
		return nil, errors.New("post not found")
	}
	views, err := s.repo.CountEvents(postID, "view")
	if err != nil {
		return nil, err
	}
	likes, err := s.repo.CountEvents(postID, "like")
	if err != nil {
		return nil, err
	}
	rawEvents, err := s.repo.RecentEvents(postID, 10)
	if err != nil {
		return nil, err
	}
	var recent []domain.AnalyticEvent
	for _, e := range rawEvents {
		recent = append(recent, domain.AnalyticEvent{
			ID:        e.ID,
			Event:     e.Event,
			UserID:    e.UserID,
			IPAddress: e.IPAddress,
			UserAgent: e.UserAgent,
			CreatedAt: e.CreatedAt,
		})
	}
	return &domain.AnalyticsSummary{
		PostID:       postID,
		ViewCount:    views,
		LikeCount:    likes,
		RecentEvents: recent,
	}, nil
}
