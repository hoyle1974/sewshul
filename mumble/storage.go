package mumble

import (

	//"syscall/js"

	"sync"

	"github.com/hoyle1974/sewshul/services"
)

type DistroStorage interface {
	Store(mdata MumbleData)
}

type distroStorage struct {
	lock    sync.Mutex
	ownerId services.AccountId
	data    []*MumbleData
}

func NewDistroStorage(ownerId services.AccountId) DistroStorage {
	return &distroStorage{ownerId: ownerId}
}

func (d *distroStorage) Store(mdata MumbleData) {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.data = append(d.data, &mdata)
}

func (d *distroStorage) Process() {
}

/*
//------------

// SessionStorage calls the native functions of the JavaScript object sessionStorage.
// For example, the following function:
//
// SessionStorage("setItem", "myItem", "itemValue")
//
// is equivalent to the JavaScript function:
//
// sessionStorage.setItem('myItem', 'itemValue')
func SessionStorage(action string, args ...interface{}) js.Value {
	return js.Global().Get("sessionStorage").Call(action, args...)

}

// LocalStorage calls the native functions of the JavaScript object localStorage.
// Example: The following function:
//
// LocalStorage("setItem", "myItem", "itemValue")
//
// is equivalent to the JavaScript function:
//
// localStorage.setItem('myItem', 'itemValue')
func LocalStorage(action string, args ...interface{}) js.Value {
	return js.Global().Get("localStorage").Call(action, args...)

}

// ObjectToStorage is capable of storing different values in the browser's storage.
// Its effectiveness has been tested with structs.
//
// The parameter storageType selects the type of storage and only accepts the values
// "localStorage" or "sessionStorage". The parameter nameItem names the item in the
// chosen storage and can be any string. The parameter object is the value
// that will be entered in the storage as a string.
//
// WARNING: Values different from "localStorage" or "sessionStorage" will trigger a panic() function.
func ObjectToStorage(storageType, nameItem string, object interface{}) {
	storage := js.Global().Get(storageType)
	data, _ := json.Marshal(object)

	storage.Call("setItem", nameItem, string(data))

}

// @Experimental
//
// ImportToStorage constructs an object from an item in the storage.
func ImportToStorage(storageType, nameItem string, v interface{}) error {
	item := js.Global().Get(storageType).Call("getItem", nameItem)
	fmt.Println(item.String())
	if !item.Truthy() {
		return errors.New("invalid storage operation")
	}

	data, err := json.Marshal(item)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)

}

*/
