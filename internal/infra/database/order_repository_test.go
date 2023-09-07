package database

import (
	"database/sql"
	"testing"

	"github.com/rodrigoafernandes/desafio-clean-architecture/internal/entity"
	"github.com/stretchr/testify/suite"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) BeforeTest(suiteName, testName string) {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) AfterTest(suiteName, testName string) {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestGivenOrdersFound_WhenSearchAllOrders_ThenShouldReturnsEntityOrdersArray() {
	firstOrder, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	secondOrder, err := entity.NewOrder("456", 10.0, 2.0)
	suite.NoError(err)
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(firstOrder)
	suite.NoError(err)
	err = repo.Save(secondOrder)
	suite.NoError(err)

	ordersFirstPage, err := repo.FindAll(1, 1, "asc")
	suite.NoError(err)
	suite.NotEmpty(ordersFirstPage)
	suite.Len(ordersFirstPage, 1)
	orderResult := ordersFirstPage[0]
	suite.Equal(firstOrder.ID, orderResult.ID)

	ordersSecondPage, err := repo.FindAll(2, 1, "cas")
	suite.NoError(err)
	suite.NotEmpty(ordersSecondPage)
	suite.Len(ordersSecondPage, 1)
	orderSecondResult := ordersSecondPage[0]
	suite.Equal(secondOrder.ID, orderSecondResult.ID)

	ordersFirstPageDesc, err := repo.FindAll(2, 1, "desc")
	suite.NoError(err)
	suite.NotEmpty(ordersFirstPageDesc)
	suite.Len(ordersFirstPageDesc, 1)
	orderFirstResultDesc := ordersFirstPageDesc[0]
	suite.Equal(firstOrder.ID, orderFirstResultDesc.ID)
}
