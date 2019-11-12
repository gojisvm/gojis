// Package snapshot is used to create and read byte dumps of objects. This
// package was created to improve startup times, for example the whole VM, which
// takes time to create. With this package, the VM can be snapshotted and
// restored to a certain state easily.
package snapshot
