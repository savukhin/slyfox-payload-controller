set (CMAKE_CXX_STANDARD 11)
cmake_minimum_required(VERSION 3.23)

project(SlyfoxPayload
    VERSION 1.0.0
    DESCRIPTION "Slyfox codefor payload project"
    LANGUAGES CXX
)

set(SRC_DIRECTORY ${PROJECT_SOURCE_DIR}/src)

add_subdirectory(${SRC_DIRECTORY})