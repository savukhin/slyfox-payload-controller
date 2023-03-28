project(lib_image_analyzer)

set(SRC
    ImageAnalyzer.cpp
)

set(MAIN_DEPENDENCIES
    ${MAIN_DEPENDENCIES}
    image_analyzer
    PARENT_SCOPE
)

add_library(image_analyzer STATIC ${SRC})
