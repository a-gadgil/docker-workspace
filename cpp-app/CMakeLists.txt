cmake_minimum_required(VERSION 3.22)
project(cpp_microservice VERSION 1.1.0 LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

add_executable(cpp_app main.cpp)

target_compile_options(cpp_app PRIVATE -Wall -Wextra -Wpedantic)

include(GNUInstallDirs)
install(TARGETS cpp_app RUNTIME DESTINATION ${CMAKE_INSTALL_BINDIR})
