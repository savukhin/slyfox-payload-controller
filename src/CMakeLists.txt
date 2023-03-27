project(hello_world)

set(
    LIBS
    Devices
    ImageAnalyzer
)
MESSAGE("Main cmake")

# set(SRC
#     main.cpp
# )

# include_directories(${LIBS})
# include_directories(${SRC_DIRECTORY})
# include_directories(Devices)
include_directories(ImageAnalyzer)
# include_directories(SensorAnalyzer)
# include_directories(MovementController)

add_subdirectory(ImageAnalyzer)
# add_subdirectory(Devices)


set(MAIN_DEPENDENCIES
    image_analyzer
)

add_subdirectory(tests)

file(GLOB SRC_FILES ${RADIO_SRC_DIRECTORY}/tests/*.cpp)

add_executable(main main.cpp)
# add_executable(main main.cpp ImageAnalyzer/ImageAnalyzer.cpp)
# target_link_libraries(main ${MAIN_DEPENDENCIES} ${SRC})
