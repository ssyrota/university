#include <bits/stdc++.h>

using namespace std;

typedef function<bool(vector<string>, vector<string>, vector<string>)> Predicate;

struct Relation : public Record
{
  Record *owner;
  Record *member;
  bool single_owner;
  bool single_member;
  vector<pair<int, int>> id_pairs;
  Relation *revese;

  string to_string()
  {
    stringstream ss;
    ss << name << '\n';
    ss << owner->name << "," << member->name << "," << single_owner << "," << single_member;
    for (auto attr : attributes)
    {
      ss << ',' << attr;
    }
    ss << '\n';
    for (int i = 0; i < id_pairs.size(); i++)
    {
      ss << id_pairs[i].first << "," << id_pairs[i].second;
      for (auto value : rows[i])
      {
        ss << "," << value;
      }
      ss << '\n';
    }
    return ss.str();
  }

  Relation filter_beta(Predicate p)
  {
    Relation result = *this;
    result.rows.clear();
    result.id_pairs.clear();

    for (int i = 0; i < id_pairs.size(); i++)
    {
      int owner_id = id_pairs[i].first;
      int member_id = id_pairs[i].second;
      if (p(owner->rows[owner_id], member->rows[member_id], rows[i]))
      {
        result.rows.push_back(rows[i]);
        result.id_pairs.push_back(id_pairs[i]);
      }
    }

    return result;
  }
};

int main()
{

  Record patients = Record::read_from_file("patients.txt");
  Record doctors = Record::read_from_file("doctors.txt");

  cout << patients.to_string() << endl;
  cout << doctors.to_string() << endl;

  Relation patient_doctor = {
      .owner = &patients,
      .member = &doctors,
      .single_owner = false,
      .single_member = false,
      .id_pairs = {{1, 1}, {1, 2}, {2, 3}, {1, 3}, {3, 4}},
      .revese = nullptr,
  };
  patient_doctor.name = "patient_doctor";
  patient_doctor.attributes = {"date"};
  patient_doctor.rows = {
      {"2019-01-01"},
      {"2019-01-02"},
      {"2019-01-03"},
      {"2019-01-04"},
      {"2019-01-05"},
  };

  cout << patient_doctor.to_string() << endl;

  cout << "INITIAL RELATION:\n";
  cout << patient_doctor.to_string() << '\n';

  cout << "FILTERED RELATION:\n";

  Relation filtered = patient_doctor.filter_beta([](auto owner, auto member, auto row)
                                                 {
        // patients with asthma
        // or doctors older than 50
        return owner[2] == "asthma" || stoi(member[3]) > 50; });

  cout << filtered.to_string() << '\n';

  ofstream file("patient_doctor.txt");
  file << filtered.to_string();

  return 0;
}