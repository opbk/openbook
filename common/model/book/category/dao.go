package category

import (
	"database/sql"
	"github.com/opbk/openbook/common/db"

	logger "github.com/cihub/seelog"
	"strconv"
	"strings"
)

const (
	FIND   = "SELECT id, parent_id, path, name, books FROM categories WHERE id = $1"
	LIST   = "SELECT id, parent_id, path, name, books FROM categories LIMIT $1 OFFSET $2"
	UPDATE = "UPDATE categories SET parent_id = $1, path = $2, name = $3, books = $4 WHERE id = $5"
	INSERT = "INSERT INTO categories (parent_id, path, name, books) VALUES ($1, $2, $3, $4) RETURNING id"
	DELETE = "DELETE FROM categories WHERE id = $1"
)

const (
	FIND_CHILD_CATEGORIES = "SELECT id, parent_id, path, name, books FROM categories WHERE parent_id = $1"
)

func connection() *sql.DB {
	return db.Connection()
}

func interateRows(rows *sql.Rows) []*Category {
	categories := make([]*Category, 0)
	for rows.Next() {
		var category *Category = new(Category)
		rows.Scan(&category.Id, &category.ParentId, &category.Path, &category.Name, &category.Books)
		categories = append(categories, category)
	}
	return categories
}

func List(limit, offset int) []*Category {
	rows, err := connection().Query(LIST, limit, offset)
	if err != nil {
		logger.Errorf("Database error while getting list of categories: %s", err)
	}

	return interateRows(rows)
}

func ListChildCategories(id int64) []*Category {
	rows, err := connection().Query(FIND_CHILD_CATEGORIES, id)
	if err != nil {
		logger.Errorf("Database error while getting list of child categories: %s", err)
	}
	return interateRows(rows)
}

func Search(filter map[string]interface{}) []*Category {
	id, ok := filter["category"]
	if ok {
		return ListChildCategories(id.(int64))
	}

	return ListChildCategories(0)
}

func Find(id int64) *Category {
	var category *Category = new(Category)
	row := connection().QueryRow(FIND, id)
	err := row.Scan(&category.Id, &category.ParentId, &category.Path, &category.Name, &category.Books)

	if err != nil {
		logger.Errorf("Database error while finding category %d: %s", id, err)
		return nil
	}

	return category
}

func (c *Category) Save() {
	if c.Id != 0 {
		c.update()
	} else {
		c.insert()
	}
}

func (c *Category) update() {
	_, err := connection().Exec(UPDATE, c.ParentId, c.Path, c.Name, c.Books, c.Id)
	if err != nil {
		logger.Errorf("Database error while updating category %d: %s", c.Id, err)
	}
}

func (c *Category) insert() {
	err := connection().QueryRow(INSERT, c.ParentId, c.Path, c.Name, c.Books).Scan(&c.Id)
	if err != nil {
		logger.Errorf("Database error while inserting category: %s", err)
	}
}

func (c *Category) Delete() {
	_, err := connection().Exec(DELETE, c.Id)
	if err != nil {
		logger.Errorf("Database error while deleting category %d: %s", c.Id, err)
	}
}

func (c *Category) GetPath() []*Category {
	var categories = make([]*Category, 0)
	if c.Path == "" {
		return categories
	}
	ids := parsePath(c.Path)
	for _, id := range ids {
		categories = append(categories, Find(id))
	}
	return categories
}

func parsePath(path string) []int64 {
	path = strings.Trim(path, ">")

	var str_cat = strings.Split(path, ">")

	var categories = make([]int64, 0)

	for _, s := range str_cat {
		c, _ := strconv.ParseInt(s, 10, 64)
		categories = append(categories, c)
	}
	return categories
}

func (c *Category) GetChildCategories() []*Category {
	return ListChildCategories(c.Id)
}
