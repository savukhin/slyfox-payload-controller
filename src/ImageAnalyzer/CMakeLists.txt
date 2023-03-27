project(lib_image_analyzer)

set(SRC
    ImageAnalyzer.cpp
)

MESSAGE("Found cmake")

# set(SRC )

add_library(image_analyzer SHARED ${SRC})

# include
