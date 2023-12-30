package handlers

import (
	"book_store/models"
	"log"

	"github.com/aymerick/raymond"
	"github.com/labstack/echo/v4"
)

type StoreBranchHandlers struct {
	Repo models.StoreBranchQuerier
}

func (h *StoreBranchHandlers) Create(c echo.Context) error {
	store := models.StoreBranch{}
	err := c.Bind(&store)
	if err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}

	store, err = h.Repo.Create(store)

	if err != nil {
		log.Println(err)
		return c.String(500, err.Error())
	}

	return c.JSON(200, store)

}

func (h *StoreBranchHandlers) List(c echo.Context) error {
	storeBranches, err := h.Repo.List()
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}

	tpl := `<h1>store branches</h1>
	{{#store_branches}}
	  <div>
		<h2>{{BranchID}}</h2>
		<p>Address: {{BranchAddress}}</p>
	  </div>
	{{/store_branches}}</div>
`

	ctx := map[string]any{
		"store_branches": storeBranches,
	}

	result, _ := raymond.Render(tpl, ctx)

	return c.HTML(200, result)

}
