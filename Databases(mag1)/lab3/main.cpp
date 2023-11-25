#include </app/relation.cpp>
#include <fstream>
using namespace std;

int main()
{
    Relation A = Relation::make_from_file("A.txt");
    std::cout << A.to_string();
    return 0;
}
