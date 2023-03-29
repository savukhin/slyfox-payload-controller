include(FetchContent)

FetchContent_Declare(
  googletest
  GIT_REPOSITORY https://github.com/google/googletest.git
  GIT_TAG        release-1.11.0
  # GIT_SHALLOW TRUE
  # GIT_PROGRESS TRUE
)
FetchContent_GetProperties(googletest)
if (NOT googletest_POPULATED)
    FetchContent_Populate(googletest)

    add_subdirectory(${googletest_SOURCE_DIR} ${googletest_BINARY_DIR})
endif ()
FetchContent_MakeAvailable(googletest)
add_library(GTest::GTest INTERFACE IMPORTED)
target_link_libraries(GTest::GTest INTERFACE gtest_main)

FILE(COPY img.png DESTINATION "${CMAKE_BINARY_DIR}/src/tests")
FILE(COPY suitable_img.png DESTINATION "${CMAKE_BINARY_DIR}/src/tests")

set(TEST_SRC 
  ImageAnalyzer.cpp
  tests.cpp
)

enable_testing ()
add_executable (Test ${TEST_SRC})
target_link_libraries(Test gtest_main gmock_main ${MAIN_DEPENDENCIES})

add_test (NAME MyTest COMMAND Test)
