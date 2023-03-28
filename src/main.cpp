#include <iostream>
#include "ImageAnalyzer/ImageAnalyzer.hpp"


int main(int argc, char** argv)
{
	std::cout << "Hello, World!" << std::endl;
	auto analyzer = slyfox::ImageAnalyzer();
	int result = slyfox::ImageAnalyzer::GetValue(1);
	std::cout << "Result is" << analyzer.GetValue(1) << std::endl;
	return 0;
}	