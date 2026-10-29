cmake_minimum_required(VERSION 3.12)
project(cpp_app)

set(CMAKE_CXX_STANDARD 17)

# Find system libraries via pkg-config
find_package(PkgConfig REQUIRED)
pkg_check_modules(PQ REQUIRED libpq)
pkg_check_modules(HIREDIS REQUIRED hiredis)

add_executable(cpp_app main.cpp)

# Include headers and link libraries
target_include_directories(cpp_app PRIVATE ${PQ_INCLUDE_DIRS} ${HIREDIS_INCLUDE_DIRS})
target_link_libraries(cpp_app PRIVATE ${PQ_LIBRARIES} ${HIREDIS_LIBRARIES} pthread)
