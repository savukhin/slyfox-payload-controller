set (CMAKE_CXX_STANDARD 23)

project(hello_world)

add_subdirectory(ImageAnalyzer)
# add_subdirectory(Devices)

add_subdirectory(tests)

add_executable(main main.cpp)
target_link_libraries(main ${MAIN_DEPENDENCIES})
