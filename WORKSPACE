http_archive(
    name = "com_google_protobuf",
    sha256 = "56b5d9e1ab2bf4f5736c4cfba9f4981fbc6976246721e7ded5602fbaee6d6869",
    strip_prefix = "protobuf-3.5.1.1",
    urls = ["https://github.com/google/protobuf/archive/v3.5.1.1.tar.gz"],
)

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/archive/1903997d0945ce92848447528718c7026b728f30.tar.gz",  # June 1, 2018
    strip_prefix = "rules_go-1903997d0945ce92848447528718c7026b728f30",
    sha256 = "c4d602a4b5204bae8e7c5cf3f916845b2cef539741d020b1e1bba09d843259b2",
)

http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/archive/644ec7202aa352b78d65bc66abc2c0616d76cc84.tar.gz",  # May 31, 2018
    strip_prefix = "bazel-gazelle-644ec7202aa352b78d65bc66abc2c0616d76cc84",
    sha256 = "ef1722f82aed4d312779e9357391b8715a0933a880d84031dde649ecac605274",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    tag = "v0.23.0",  # May 18, 2018
)

go_repository(
    name = "org_golang_google_api",
    commit = "00e3bb8d04691e25ee2fccf98c866bcb7925c3ec",  # June 5, 2018
    importpath = "google.golang.org/api",
)

go_repository(
    name = "org_golang_x_sync",
    commit = "1d60e4601c6fd243af51cc01ddf169918a5407ca",  # March 14, 2018
    importpath = "golang.org/x/sync",
)

go_repository(
    name = "org_golang_x_oauth2",
    commit = "cdc340f7c179dbbfa4afd43b7614e8fcadde4269",  # May 2, 2018
    importpath = "golang.org/x/oauth2",
)

go_repository(
    name = "io_opencensus_go",
    importpath = "go.opencensus.io",
    tag = "v0.12.0",  # June 5, 2018
)

go_repository(
    name = "com_github_googleapis_gax_go",
    importpath = "github.com/googleapis/gax-go",
    tag = "v2.0.0",  # September 14, 2017
)

new_http_archive(
    name = "googleapis_archive",
    build_file = "//third_party:googleapis.BUILD",
    sha256 = "df585427959ebb9949ca12a45888b50a5b8edcb2b270aea1800225d159a49ff6",
    strip_prefix = "googleapis-3273178ea4684acc4f512f7bef7349dd72db88f6",
    urls = ["https://github.com/googleapis/googleapis/archive/3273178ea4684acc4f512f7bef7349dd72db88f6.tar.gz"],
)
