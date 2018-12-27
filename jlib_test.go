package jlib

import (
	"reflect"
	"testing"
)

var obj_data = Object{
	"src":  "Images/Sun.png",
	"name": "sun1",
}

var arr_data = Array{
	"play game", "watch movie", 10, 20.25,
}

var test_data = NewObject(Object{
	"id":    1,
	"money": 1.0,
	"name":  "ATTHDEV",
	"hobby": arr_data,
	"image": obj_data,
})

func TestNewObject(t *testing.T) {
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name string
		args args
		want *JSONObject
	}{
		{"",
			args{[]interface{}{}},
			NewObject(),
		},
		{"",
			args{[]interface{}{[]byte(`{"Id":"1234"}`)}},
			NewObject(`{"Id":"1234"}`),
		},
		{"",
			args{[]interface{}{`{"Id":"1234"}`}},
			NewObject(`{"Id":"1234"}`),
		},
		{"",
			args{[]interface{}{Object{"Id": "1234"}}},
			NewObject(Object{"Id": "1234"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewObject(tt.args.i...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewArray(t *testing.T) {
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name string
		args args
		want *JSONArray
	}{
		// TODO: Add test cases.
		{"",
			args{[]interface{}{}},
			NewArray(),
		},
		{"",
			args{[]interface{}{[]byte(`["Home", "Office"]`)}},
			NewArray(`["Home", "Office"]`),
		},
		{"",
			args{[]interface{}{`["Home", "Office"]`}},
			NewArray(`["Home", "Office"]`),
		},
		{"",
			args{[]interface{}{Array{"Home", "Office"}}},
			NewArray(Array{"Home", "Office"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArray(tt.args.i...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONObject_String(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		j    *JSONObject
		args args
		want string
	}{
		// TODO: Add test cases.
		{"", test_data, args{"name"}, "ATTHDEV"},
		{"", NewObject(), args{"name"}, ""},
		{"", test_data, args{"na_me"}, ""},
		{"", test_data, args{"id"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.String(tt.args.key); got != tt.want {
				t.Errorf("JSONObject.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONObject_Int(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		j    *JSONObject
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", test_data, args{"id"}, 1},
		{"", test_data, args{"money"}, 1},
		{"", NewObject(), args{"id"}, 0},
		{"", test_data, args{"na_me"}, 0},
		{"", test_data, args{"name"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Int(tt.args.key); got != tt.want {
				t.Errorf("JSONObject.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONObject_Float(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		j    *JSONObject
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"", test_data, args{"id"}, 1.0},
		{"", test_data, args{"money"}, 1},
		{"", NewObject(), args{"money"}, 0},
		{"", test_data, args{"na_me"}, 0},
		{"", test_data, args{"name"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Float(tt.args.key); got != tt.want {
				t.Errorf("JSONObject.Float() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONObject_Bool(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		j    *JSONObject
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", NewObject(Object{"isReady": true}), args{"isReady"}, true},
		{"", NewObject(Object{"isReady": true}), args{"is__Ready"}, false},
		{"", NewObject(), args{"isReady"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Bool(tt.args.key); got != tt.want {
				t.Errorf("JSONObject.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONObject_Object(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		j    *JSONObject
		args args
		want *JSONObject
	}{
		// TODO: Add test cases.
		{"", test_data, args{"image"}, NewObject(obj_data)},
		{"", NewObject(), args{"image"}, &JSONObject{}},
		{"", test_data, args{"na_me"}, &JSONObject{}},
		{"", test_data, args{"name"}, &JSONObject{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Object(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONObject.Object() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONObject_Array(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		j    *JSONObject
		args args
		want *JSONArray
	}{
		// TODO: Add test cases.
		{"", test_data, args{"hobby"}, NewArray(arr_data)},
		{"", NewObject(), args{"hobby"}, &JSONArray{}},
		{"", test_data, args{"na_me"}, &JSONArray{}},
		{"", test_data, args{"name"}, &JSONArray{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Array(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONObject.Array() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONObject_Size(t *testing.T) {
	tests := []struct {
		name string
		j    *JSONObject
		want int
	}{
		// TODO: Add test cases.
		{"", test_data, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Size(); got != tt.want {
				t.Errorf("JSONObject.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONArray_Size(t *testing.T) {
	tests := []struct {
		name string
		j    *JSONArray
		want int
	}{
		// TODO: Add test cases.
		{"", NewArray(arr_data), 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Size(); got != tt.want {
				t.Errorf("JSONArray.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONObject_Error(t *testing.T) {
	tests := []struct {
		name    string
		j       *JSONObject
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", NewObject(obj_data), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.j.Error(); (err != nil) != tt.wantErr {
				t.Errorf("JSONObject.Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJSONArray_Error(t *testing.T) {
	tests := []struct {
		name    string
		j       *JSONArray
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", NewArray(arr_data), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.j.Error(); (err != nil) != tt.wantErr {
				t.Errorf("JSONArray.Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJSONObject_Put(t *testing.T) {
	var before = NewObject(Object{
		"id":    1,
		"money": 1.0,
		"name":  "ATTHDEV",
		"image": obj_data,
	})

	type args struct {
		key string
		v   interface{}
	}
	tests := []struct {
		name string
		j    *JSONObject
		args args
		want *JSONObject
	}{
		// TODO: Add test cases.
		{"", before, args{"hobby", arr_data}, before.Put("hobby", arr_data)},
		{"", NewObject(), args{"hobby", arr_data}, NewObject().Put("hobby", arr_data)},
		{"", before, args{"hobby", NewArray(arr_data)}, before.Put("values", NewArray(arr_data))},
		{"", before, args{"hobby", NewObject(obj_data)}, before.Put("hobby", NewObject(obj_data))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Put(tt.args.key, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONObject.Put() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONObject_Delete(t *testing.T) {
	type args struct {
		key []string
	}
	tests := []struct {
		name string
		j    *JSONObject
		args args
		want *JSONObject
	}{
		// TODO: Add test cases.
		{"", test_data, args{[]string{"money", "hobby"}}, test_data.Delete("money", "hobby")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Delete(tt.args.key...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONObject.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONObject_ToString(t *testing.T) {
	data := NewObject(obj_data)
	tests := []struct {
		name    string
		j       *JSONObject
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", data, `{"name":"sun1","src":"Images/Sun.png"}`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.ToString()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONObject.ToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JSONObject.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONArray_ToString(t *testing.T) {
	data := NewArray(arr_data)
	tests := []struct {
		name    string
		j       *JSONArray
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", data, `["play game","watch movie",10,20.25]`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.ToString()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONArray.ToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JSONArray.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONObject_ToPrettyString(t *testing.T) {
	data := NewObject(obj_data)
	resulte := `{
	"name": "sun1",
	"src": "Images/Sun.png"
}`
	tests := []struct {
		name    string
		j       *JSONObject
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", data, resulte, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.ToPrettyString()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONObject.ToPrettyString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JSONObject.ToPrettyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONArray_ToPrettyString(t *testing.T) {
	data := NewArray(arr_data)
	resulte := `[
	"play game",
	"watch movie",
	10,
	20.25
]`
	tests := []struct {
		name    string
		j       *JSONArray
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", data, resulte, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.ToPrettyString()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONArray.ToPrettyString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JSONArray.ToPrettyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONArray_String(t *testing.T) {
	data := NewArray(arr_data)
	type args struct {
		index int
	}
	tests := []struct {
		name string
		j    *JSONArray
		args args
		want string
	}{
		// TODO: Add test cases.
		{"", data, args{10}, ""},
		{"", data, args{0}, "play game"},
		{"", data, args{2}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.String(tt.args.index); got != tt.want {
				t.Errorf("JSONArray.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONArray_Int(t *testing.T) {
	data := NewArray(arr_data)
	type args struct {
		index int
	}
	tests := []struct {
		name string
		j    *JSONArray
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", data, args{10}, 0},
		{"", data, args{2}, 10},
		{"", data, args{3}, 20},
		{"", data, args{0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Int(tt.args.index); got != tt.want {
				t.Errorf("JSONArray.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONArray_Float(t *testing.T) {
	data := NewArray(arr_data)
	type args struct {
		index int
	}
	tests := []struct {
		name string
		j    *JSONArray
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"", data, args{10}, 0},
		{"", data, args{2}, 10.0},
		{"", data, args{3}, 20.25},
		{"", data, args{0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Float(tt.args.index); got != tt.want {
				t.Errorf("JSONArray.Float() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONArray_Bool(t *testing.T) {
	data := NewArray(Array{
		"ABCD", true,
	})
	type args struct {
		index int
	}
	tests := []struct {
		name string
		j    *JSONArray
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", data, args{10}, false},
		{"", data, args{0}, false},
		{"", data, args{1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Bool(tt.args.index); got != tt.want {
				t.Errorf("JSONArray.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONArray_Object(t *testing.T) {
	data := NewArray(Array{
		"ABCD", obj_data,
	})
	type args struct {
		index int
	}
	tests := []struct {
		name string
		j    *JSONArray
		args args
		want *JSONObject
	}{
		// TODO: Add test cases.
		{"", data, args{10}, &JSONObject{}},
		{"", data, args{0}, &JSONObject{}},
		{"", data, args{1}, NewObject(obj_data)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Object(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONArray.Object() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONArray_Array(t *testing.T) {
	data := NewArray(Array{
		obj_data, arr_data,
	})
	type args struct {
		index int
	}
	tests := []struct {
		name string
		j    *JSONArray
		args args
		want *JSONArray
	}{
		// TODO: Add test cases.
		{"", data, args{10}, &JSONArray{}},
		{"", data, args{0}, &JSONArray{}},
		{"", data, args{1}, NewArray(arr_data)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Array(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONArray.Array() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONArray_Add(t *testing.T) {
	data := NewArray(Array{
		1, 2,
	})
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		j    *JSONArray
		args args
		want *JSONArray
	}{
		// TODO: Add test cases.
		{"", data, args{3}, NewArray(Array{1, 2, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Add(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONArray.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONArray_DeleteAt(t *testing.T) {
	data := NewArray(Array{
		1, 2, 3,
	})
	type args struct {
		i int
	}
	tests := []struct {
		name string
		j    *JSONArray
		args args
		want *JSONArray
	}{
		// TODO: Add test cases.
		{"", data, args{1}, NewArray(Array{1, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.DeleteAt(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONArray.DeleteAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONObject_ToMap(t *testing.T) {
	tests := []struct {
		name string
		j    *JSONObject
		want Object
	}{
		// TODO: Add test cases.
		{"", test_data, test_data.data},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.ToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONObject.ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONArray_ToArray(t *testing.T) {
	test_data := NewArray(Array{
		1, 2, 3,
	})
	tests := []struct {
		name string
		j    *JSONArray
		want Array
	}{
		// TODO: Add test cases.
		{"", test_data, test_data.data},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONArray.ToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
