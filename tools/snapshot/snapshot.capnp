using Go = import "/go.capnp";
@0xbf39c8350769d85a;
$Go.package("snapshot");
$Go.import("tools/snapshot");

struct Snapshot {
    uintptrsize @0 :Int64;
    signature @1 :Data;

    nested @2 :Nested;
}

struct Nested {
    data @0 :Data;
    pointers @1 :List(Pointer);
}

struct Pointer {
    offset @0 :Int64;
    target @1 :Nested;
}