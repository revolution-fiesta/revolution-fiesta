cmake_minimum_required(VERSION 3.25)
project(Template VERSION 0.0.1)

find_package(fmt CONFIG REQUIRED)

file(GLOB SOURCES "src/*.cc" "src/*.cpp")
file(GLOB INCLUDE_DIR "include")

add_executable(Main ${SOURCES})

target_include_directories(Main PRIVATE ${INCLUDE_DIR})

target_link_libraries(Main
  fmt::fmt
)