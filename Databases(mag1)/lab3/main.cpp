#include </app/relation.cpp>
#include <fstream>
using namespace std;

int main()
{
    Relation A = Relation::make_from_file("A.txt");
    Relation B = Relation::make_from_file("B.txt");
    std::cout << A.to_string();
    std::cout << B.to_string();
    return 0;
}
