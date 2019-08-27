package suite

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
	OldVariableThatShouldStartAtFive int
}

func (suite *ExampleTestSuite) SetupTest() {
	suite.OldVariableThatShouldStartAtFive = VariableThatShouldStartAtFive
	VariableThatShouldStartAtFive = 5
}

func (suite *ExampleTestSuite) TestExample() {
	assert.Equal(suite.T(), 5, VariableThatShouldStartAtFive)
	suite.Equal(5, VariableThatShouldStartAtFive)
}

func (suite *ExampleTestSuite) TearDownTest() {
	// 进行资源释放或者数据还原
	VariableThatShouldStartAtFive = suite.OldVariableThatShouldStartAtFive
}

func (suite *ExampleTestSuite) BeforeTest() {
	// 测试前置处理
}

func (suite *ExampleTestSuite) AfterTest() {
	// 测试后置处理
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))

	log.Println("全局变量为: ", VariableThatShouldStartAtFive)
}
