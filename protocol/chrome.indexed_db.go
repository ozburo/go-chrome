package protocol

import (
	indexedDB "github.com/mkenney/go-chrome/protocol/indexed_db"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
IndexedDB provides a namespace for the Chrome IndexedDB protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/ EXPERIMENTAL.
*/
var IndexedDB = IndexedDBProtocol{}

/*
IndexedDBProtocol defines the IndexedDB protocol methods.
*/
type IndexedDBProtocol struct{}

/*
ClearObjectStore clears all entries from an object store.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-clearObjectStore
*/
func (IndexedDBProtocol) ClearObjectStore(
	socket sock.Socketer,
	params *indexedDB.ClearObjectStoreParams,
) error {
	command := sock.NewCommand("IndexedDB.clearObjectStore", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
DeleteDatabase deletes a database.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteDatabase
*/
func (IndexedDBProtocol) DeleteDatabase(
	socket sock.Socketer,
	params *indexedDB.DeleteDatabaseParams,
) error {
	command := sock.NewCommand("IndexedDB.deleteDatabase", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
DeleteObjectStoreEntries deletes a range of entries from an object store.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteObjectStoreEntries
*/
func (IndexedDBProtocol) DeleteObjectStoreEntries(
	socket sock.Socketer,
	params *indexedDB.DeleteObjectStoreEntriesParams,
) error {
	command := sock.NewCommand("IndexedDB.deleteObjectStoreEntries", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
Disable disables events from backend.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-disable
*/
func (IndexedDBProtocol) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("IndexedDB.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables events from backend.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-enable
*/
func (IndexedDBProtocol) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("IndexedDB.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
RequestData requests data from object store or index.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestData
*/
func (IndexedDBProtocol) RequestData(
	socket sock.Socketer,
	params *indexedDB.RequestDataParams,
) (*indexedDB.RequestDataResult, error) {
	command := sock.NewCommand("IndexedDB.requestData", params)
	result := &indexedDB.RequestDataResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
RequestDatabase requests database with given name in given frame.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabase
*/
func (IndexedDBProtocol) RequestDatabase(
	socket sock.Socketer,
	params *indexedDB.RequestDatabaseParams,
) (*indexedDB.RequestDatabaseResult, error) {
	command := sock.NewCommand("IndexedDB.requestDatabase", params)
	result := &indexedDB.RequestDatabaseResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
RequestDatabaseNames requests database names for given security origin.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabaseNames
*/
func (IndexedDBProtocol) RequestDatabaseNames(
	socket sock.Socketer,
	params *indexedDB.RequestDatabaseNamesParams,
) (*indexedDB.RequestDatabaseNamesResult, error) {
	command := sock.NewCommand("IndexedDB.requestDatabaseNames", params)
	result := &indexedDB.RequestDatabaseNamesResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}
