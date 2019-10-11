# Versioning
This article describes the versioning and release process for gojis.

## Versions
We use different version patterns for different kinds of releases.

```
dev preview -> alpha -> beta -> release candidate -> release
```

### Developer previews
Dev previews are representes as tags only (no GitHub releases).
The naming scheme is `v0.0.1-dp1`, which would be the first dev preview for the upcoming release `0.0.1`.

There can be multiple dev preview versions for the same release, as in `v0.0.1-dp1`, `v0.0.1-dep2` etc. However, developer previews are no GitHub release and therefor not intended for public use.

### Alpha versions
Alpha versions are GitHub releases, intended for public use, but not recommended for prod.
The naming scheme is `v0.0.1-alpha`, which is the alpha version for the upcoming release `0.0.1`.

Between alpha and beta, APIs can change, and features might be added and/or removed.

### Beta versions
Beta versions are GitHub releases, intended for public use, but not recommended for prod.
The naming scheme is `v0.0.1-beta`, which is the beta version for the upcoming release `0.0.1`.

Beta versions are intended for public testing, which means, that the documentation has to be complete and the API has to be stable until the actual release.

### Release candidates
Release candidates are GitHub releases, intended for public use.
The naming scheme is `v1.0.0-rc1`, which is the first release candidate for the upcoming release `1.0.0`.

Release candidates are intended for public use, which means, that the documentation has to be complete and the API has to be stable until the actual release.

Release candidates are developed on a separate branch `release/1.0.0`. On this branch, only bug fixes must be merged. Once the release candidate is considered stable enough, the branch will be merged back into develop and master. The actual release is then created from the master branch, including all fixes to the release candidate.

### Releases
Releases are Github releases, intended for public use and prod ready.
The naming scheme is `v1.0.0`, which is the public release of the version `1.0.0`.

Releases must be created from the master branch.