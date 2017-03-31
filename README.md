# go-validate

各种常见的数据校验方法，所有的方法返回值都是bool，会对原始数据进行强效的数据转换。

Muti-kind of validate funcation, all of the function returns bool, will try best to convert value.

``` go

// Example
b1 := "true"
b2 := true
b3 := "T"
fmt.Println(validate.CheckBool(b1),validate.CheckBool(b2),validate.CheckBool(b3)) // true,true,true

```

## CheckType

## CheckIntRange

## CheckFloat64Range

## CheckRegexp

## CheckEmail

## CheckMobile

## CheckIPv4

## CheckRealNumber

## CheckLen

## CheckMin

## CheckMax

## CheckMaxSize

## CheckMinSize

## IsValidBoolean

## IsValidNumber

## IsArray