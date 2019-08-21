using Go = import "/go.capnp";
@0xbf39c8350769d85a;
$Go.package("snapshot");
$Go.import("tools/snapshot");

struct Snapshot {
    nested @0 :Nested;
}

struct Nested {
    data @0 :Data;
    len @1 :Int64;
    pointers @2 :List(Pointer);
}

struct Pointer {
    offset @0 :Int64;
    target @1 :Nested;
}