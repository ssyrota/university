#include </app/datasource.cpp>
#include <fstream>
using namespace std;
// 25) ВFilter(ім’яDS1, умова) -&gt; ім’яDS2 (це ім’я слід запам’ятати в сист.каталозі);

bool filter(vector<string> owner, vector<string> member, vector<string> row)
{
    return owner[0] == "test1" || owner[0] == "test3";
}

int main()
{
    Relation A = Relation::make_from_file("A.txt");
    Relation B = Relation::make_from_file("B.txt");
    DataSource A_B = DataSource::new_default(A, B);

    cout << A.to_string() << "\n-----\n";
    cout << B.to_string() << "\n-----\n";
    cout << A_B.to_string() << "\n-----\n";
    A_B.to_file();

    cout << "starting filtering"
         << "\n-----\n";
    DataSource filtered = A_B.filter_beta(filter);
    cout << "finishing filtering"
         << "\n-----\n";

    filtered.name = "filtered";
    filtered.to_file();
    cout << filtered.to_string();
    return 0;
}
