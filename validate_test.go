package validate

import (
	"testing"

	"math/rand"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckRegex(t *testing.T) {
	//正则测试以通过Email、Mobile测试为准
}

func TestCheckMobile(t *testing.T) {
	//普通手机号一组

	Convey("测试普通手机号", t, func() {
		So(CheckMobile("136547896324"), ShouldBeTrue)
		So(CheckMobile("+86136547896324"), ShouldBeTrue)
		So(CheckMobile("86136547896324"), ShouldBeTrue)
	})

	Convey("测试非法手机号", t, func() {
		So(CheckMobile("helloWorld"), ShouldBeFalse)
		So(CheckMobile("发了什么事"), ShouldBeFalse)
		So(CheckMobile("8613647896"), ShouldBeFalse)
		So(CheckMobile("+86136547324"), ShouldBeFalse)
		So(CheckMobile(13225533665), ShouldBeTrue)
	})
}

func TestCheckType(t *testing.T) {
	Convey("测试string类型", t, func() {
		So(CheckType("value interface{}", "string"), ShouldBeTrue)
		So(CheckType(398, "string"), ShouldBeFalse)
	})
	Convey("测试int类型", t, func() {
		So(CheckType(123, "int"), ShouldBeTrue)
		So(CheckType("value interface{}", "int"), ShouldBeFalse)
	})
	Convey("测试float类型", t, func() {
		So(CheckType(12.3, "float"), ShouldBeTrue)
		So(CheckType(123, "float"), ShouldBeFalse)
		So(CheckType("value interface{}", "float"), ShouldBeFalse)
	})
}

func TestCheckIntRange(t *testing.T) {
	min := rand.Int()
	max := rand.Int()
	if min > max {
		temp := min
		min = max
		max = temp
	}
	Convey("测试随机区间", t, func() {
		rand := rand.Int()
		So(CheckIntRange(rand, min, max), ShouldEqual, min <= rand && rand <= max)
	})
	Convey("测试负正区间", t, func() {
		So(CheckIntRange(0, -1, 1), ShouldBeTrue)
		So(CheckIntRange(-10, -1, 1), ShouldBeFalse)
		So(CheckIntRange(10, -1, 1), ShouldBeFalse)
	})
	Convey("测试负负区间", t, func() {
		So(CheckIntRange(-10, -15, -5), ShouldBeTrue)
		So(CheckIntRange(0, -15, -1), ShouldBeFalse)
		So(CheckIntRange(10, -15, -1), ShouldBeFalse)
	})
	Convey("测试正正区间", t, func() {
		So(CheckIntRange(5, 1, 10), ShouldBeTrue)
		So(CheckIntRange(-10, 1, 10), ShouldBeFalse)
		So(CheckIntRange(15, 1, 10), ShouldBeFalse)
	})
}

func TestFloatRange(t *testing.T) {
	//有问题，待定
	min := rand.Float64()
	max := rand.Float64()
	if min > max {
		temp := min
		min = max
		max = temp
	}
	Convey("测试随机区间", t, func() {
		rand := rand.Float64()
		So(CheckFloat64Range(rand, min, max), ShouldEqual, min <= rand && rand <= max)
	})
	Convey("测试负正区间", t, func() {
		So(CheckFloat64Range(0., -1.0, 1.0), ShouldBeTrue)
		So(CheckFloat64Range(-10., -1., 1.), ShouldBeFalse)
		So(CheckFloat64Range(10., -1., 1.), ShouldBeFalse)
	})
	Convey("测试负负区间", t, func() {
		So(CheckFloat64Range(-10., -15., -5.), ShouldBeTrue)
		So(CheckFloat64Range(0., -15., -1.), ShouldBeFalse)
		So(CheckFloat64Range(10., -15., -1.), ShouldBeFalse)
	})
	Convey("测试正正区间", t, func() {
		So(CheckFloat64Range(5., 1., 10.), ShouldBeTrue)
		So(CheckFloat64Range(-10., 1., 10.), ShouldBeFalse)
		So(CheckFloat64Range(15., 1., 10.), ShouldBeFalse)
	})
}

func TestCheckEmail(t *testing.T) {
	Convey("测试email", t, func() {
		So(CheckEmail("value interface{}"), ShouldBeFalse)
		So(CheckEmail("haha eee@fff.c"), ShouldBeFalse)
		So(CheckEmail("haha.eee@fff.c"), ShouldBeFalse)

		So(CheckEmail("haha.eee@fff.cc"), ShouldBeTrue)
	})
}

func TestCheckIPv4(t *testing.T) {
	Convey("测试IPv4", t, func() {
		So(CheckIPv4("value interface{}"), ShouldBeFalse)
		So(CheckIPv4("1257.26.34.1111"), ShouldBeFalse)

		So(CheckIPv4("120.25.82.60"), ShouldBeTrue)
		So(CheckIPv4("0.015.250.220"), ShouldBeTrue)
	})
}

func TestCheckLen(t *testing.T) {
	Convey("测试slice、array、map或string的len", t, func() {
		Convey("测试Len(string)", func() {
			So(CheckLen("value interface{}", 17), ShouldBeTrue)
		})

		Convey("测试Len(array)", func() {

			testAry := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
			So(CheckLen(testAry, 10), ShouldBeTrue)
		})

		Convey("测试Len(slice)", func() {
			testSlice := make([]int, 18)
			So(CheckLen(testSlice, 18), ShouldBeTrue)
		})

		Convey("测试Len(map)", func() {
			testMap := make(map[int]int)
			testMap[1] = 1
			So(CheckLen(testMap, 1), ShouldBeTrue)
		})

	})
}

func TestCheckMin(t *testing.T) {
	Convey("测试string、int、float、uint的min", t, func() {
		Convey("测试string的min(肯定测试)", func() {
			So(CheckMin("5", 4), ShouldBeTrue)
		})

		Convey("测试string的min(否定测试)", func() {
			So(CheckMin("3", 4), ShouldBeFalse)
		})

		Convey("测试int的min(肯定测试)", func() {
			So(CheckMin(5, 4), ShouldBeTrue)
		})

		Convey("测试int的min(否定测试)", func() {
			So(CheckMin(3, 4), ShouldBeFalse)
		})

		Convey("测试float的min(肯定测试)", func() {
			So(CheckMin(5.1, 4), ShouldBeTrue)
		})

		Convey("测试float的min(否定测试)", func() {
			So(CheckMin(3.1, 4), ShouldBeFalse)
		})

	})

}

func TestCheckMax(t *testing.T) {
	Convey("测试string、int、float、uint的max", t, func() {
		Convey("测试string的max(肯定测试)", func() {
			So(CheckMax("5", 10), ShouldBeTrue)
		})

		Convey("测试string的max(否定测试)", func() {
			So(CheckMax("5", 4), ShouldBeFalse)
		})

		Convey("测试int的max(肯定测试)", func() {
			So(CheckMax(5, 10), ShouldBeTrue)
		})

		Convey("测试int的max(否定测试)", func() {
			So(CheckMax(5, 4), ShouldBeFalse)
		})

		Convey("测试float的max(肯定测试)", func() {
			So(CheckMax(5.1, 10), ShouldBeTrue)
		})

		Convey("测试float的max(否定测试)", func() {
			So(CheckMax(5.1, 4), ShouldBeFalse)
		})

	})

}
