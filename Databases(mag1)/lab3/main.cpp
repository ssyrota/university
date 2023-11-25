#include </app/relation.cpp>
#include <fstream>
using namespace std;

int main()
{
    Relation A = Relation::make_from_file("A.txt");
    Relation B = Relation::make_from_file("B.txt");
    A.to_file();
    B.to_file();
    return 0;
}
