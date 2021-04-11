package graph

import "github.com/MrChampz/fullcycle-comunicacao/graphql/graph/model"

type Resolver struct {
	Categories []*model.Category
	Courses    []*model.Course
	Chapters   []*model.Chapter
}
