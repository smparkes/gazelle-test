load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://github.com/bazelbuild/rules_go/archive/5e69779e3fe39126bf4fc14912bc766ada10efec/5e69779e3fe39126bf4fc14912bc766ada10efec.tar.gz",
    ],
    strip_prefix = "rules_go-5e69779e3fe39126bf4fc14912bc766ada10efec",
    sha256 = "4c05784851b57f78765999049fd094993100b196d0100bb36a142568213764f1",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

http_archive(
    name = "bazel_gazelle",
    urls = [
        "https://github.com/bazelbuild/bazel-gazelle/archive/33ccc50dcd2d0b9d8937bc8497fa13505bd7a665/bazel-gazelle-33ccc50dcd2d0b9d8937bc8497fa13505bd7a665.tar.gz",
    ],
    strip_prefix = "bazel-gazelle-33ccc50dcd2d0b9d8937bc8497fa13505bd7a665",
    sha256 = "a2e9a3006b3277b6ba284f89334a41699276dcaca24703c0638cd79309e1c9d1",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

http_archive(
    name = "rules_proto",
    sha256 = "602e7161d9195e50246177e7c55b2f39950a9cf7366f74ed5f22fd45750cd208",
    strip_prefix = "rules_proto-97d8af4dc474595af3900dd85cb3a29ad28cc313",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()
