#include </app/datasource.cpp>
#include <fstream>
using namespace std;

// 25) ВFilter(ім’яDS1, умова) -&gt; ім’яDS2 (це ім’я слід запам’ятати в сист.каталозі);
int main()
{
    Relation A = Relation::make_from_file("A.txt");
    Relation B = Relation::make_from_file("B.txt");
    DataSource A_B = DataSource::new_default(A, B);
    A_B.to_file();
    return 0;
}
