package dto

type InputPosts struct {
	Title   *string `json:"title,omitempty" doc:"Title of the posts"`
	Content *string `json:"content,omitempty" doc:"Content of the posts"`
	UserID  *string `json:"user_id,omitempty" doc:"User ID of the posts"`
}
