package kgo

import (
	"reflect"
	"testing"
)

func TestJoinStructsField(t *testing.T) {
	type User struct {
		Id    int
		Name  string
		Score float64
	}
	var users = []User{
		{
			Id:    1,
			Name:  "one",
			Score: 1.1,
		},
		{
			Id:    2,
			Name:  "two",
			Score: 2.2,
		},
		{
			Id:    3,
			Name:  "three",
			Score: 3.3,
		},
	}
	type Args struct {
		name           string
		elements       []User
		usingFieldName string
		excepted       string
	}
	testCases := []Args{
		{
			name:           "JoinStructsIntFieldCase",
			elements:       users,
			usingFieldName: "Id",
			excepted:       "1,2,3",
		},
		{
			name:           "JoinStructsStringFieldCase",
			elements:       users,
			usingFieldName: "Name",
			excepted:       "one,two,three",
		},
		{
			name:           "JoinStructsFloatFieldCase",
			elements:       users,
			usingFieldName: "Score",
			excepted:       "1.100000,2.200000,3.300000", // 浮点型默认有6位小数位，实际使用中不会这么判断，得到值就行了
		},
	}
	for _, testCase := range testCases {
		if actual := JoinStructsField(testCase.elements, testCase.usingFieldName); actual != testCase.excepted {
			t.Errorf("JoinStructsField，输入：%+v，预期：%+v，实际：%+v", testCase.elements, testCase.excepted, actual)
		}
	}
}

func TestPickStructsField(t *testing.T) {
	type User struct {
		Id    int
		Name  string
		Score float64
	}
	var users = []User{
		{
			Id:    1,
			Name:  "one",
			Score: 1.1,
		},
		{
			Id:    2,
			Name:  "two",
			Score: 2.2,
		},
		{
			Id:    3,
			Name:  "tree",
			Score: 3.3,
		},
	}
	type TestCase struct {
		name           string
		elements       []User
		usingFieldName string
		exceptedId     []int
		exceptedName   []string
		exceptedScore  []float64
	}
	testCases := []TestCase{
		{
			name:           "PickStructsIntFieldCase",
			elements:       users,
			usingFieldName: "id",
			exceptedId:     []int{1, 2, 3},
		},
		{
			name:           "PickStructsStringFieldCase",
			elements:       users,
			usingFieldName: "name",
			exceptedName:   []string{"one,two,three"},
		},
		{
			name:           "JoinStructsFloatFieldCase",
			elements:       users,
			usingFieldName: "score",
			exceptedScore:  []float64{1.1, 2.2, 3.3},
		},
	}
	for _, testCase := range testCases {
		switch testCase.usingFieldName {
		case "id":
			if actual := PickStructsField[User, int](testCase.elements, testCase.usingFieldName); reflect.DeepEqual(actual, testCase.exceptedId) {
				t.Errorf("字符串清理，输入：%+v，预期：%+v，实际：%+v", testCase.elements, testCase.exceptedId, actual)
			}
			break
		case "name":
			if actual := PickStructsField[User, string](testCase.elements, testCase.usingFieldName); reflect.DeepEqual(actual, testCase.exceptedName) {
				t.Errorf("字符串清理，输入：%+v，预期：%+v，实际：%+v", testCase.elements, testCase.exceptedName, actual)
			}
			break
		case "score":
			if actual := PickStructsField[User, float64](testCase.elements, testCase.usingFieldName); reflect.DeepEqual(actual, testCase.exceptedScore) {
				t.Errorf("字符串清理，输入：%+v，预期：%+v，实际：%+v", testCase.elements, testCase.exceptedScore, actual)
			}
			break
		default:
			break
		}
	}
}
