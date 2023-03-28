cmake_minimum_required(VERSION 3.23)

project(lib_image_analyzer)

set(SRC
    ImageAnalyzer.cpp
)

set(MAIN_DEPENDENCIES
    ${MAIN_DEPENDENCIES}
    image_analyzer
    PARENT_SCOPE
)

find_package( OpenCV REQUIRED )
include_directories(${OpenCV_INCLUDE_DIRS})
# FILE(COPY img.png DESTINATION "${CMAKE_BINARY_DIR}/src/ImageAnalyzer")

# file(GLOB_RECURSE SOURCES RELATIVE ${CMAKE_SOURCE_DIR} "*.png")
# add_executable(main ${SOURCES})

# add_library(image_analyzer STATIC ${SOURCES})
add_library(image_analyzer STATIC ${SRC})
target_link_libraries( image_analyzer ${OpenCV_LIBS} )
