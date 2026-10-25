// swift-tools-version: 5.9
import PackageDescription

let package = Package(
    name: "swift-app",
    platforms: [
        .macOS(.v13)
    ],
    products: [
        .executable(name: "swift-app", targets: ["swift-app"])
    ],
    targets: [
        .executableTarget(
            name: "swift-app",
            path: "Sources"
        )
    ]
)
