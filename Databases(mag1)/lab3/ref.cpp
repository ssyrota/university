#include <bits/stdc++.h>
#include <iostream>
// 25) ВFilter(ім’яDS1, умова) -&gt; ім’яDS2 (це ім’я слід запам’ятати в сист.каталозі);

using namespace std;

struct Relation
{
  string name;
  vector<string> attributes;
  vector<vector<string>> rows;

  static Relation read_from_file(const char *file_name)
  {
    ifstream file(file_name);
    string line;
    getline(file, line);
    stringstream ss(line);
    string token;
    getline(ss, token);

    Relation record;
    record.name = token;
    while (getline(ss, token, ','))
    {
      record.attributes.push_back(token);
    }

    while (getline(file, line))
    {
      stringstream ss(line);
      string token;
      vector<string> row;
      while (getline(ss, token, ','))
      {
        row.push_back(token);
      }
      record.rows.push_back(row);
    }
    return record;
  }

  string to_string()
  {
    stringstream ss;
    ss << name;
    for (auto attribute : attributes)
    {
      ss << attribute << ",";
    }
    ss << '\n';
    for (auto row : rows)
    {
      for (auto value : row)
      {
        ss << value << ",";
      }
      ss << '\n';
    }
    return ss.str();
  }
};

typedef function<bool(vector<string>, vector<string>, vector<string>)> Predicate;

struct DS : public Relation
{
  Relation *owner;
  Relation *member;
  bool single_owner;
  bool single_member;
  vector<pair<int, int>> id_pairs;
  DS *revese;

  static DS make(Relation A, Relation B)
  {
    DS A_B = {
        .owner = &A,
        .member = &B,
        .single_owner = false,
        .single_member = false,
        .id_pairs = {{1, 1}, {1, 2}, {2, 3}, {1, 3}, {3, 4}},
        .revese = nullptr,
    };
    A_B.name = "A_B";
    A_B.attributes = {"date"};
    A_B.rows = {
        {"2019-01-01"},
        {"2019-01-02"},
        {"2019-01-03"},
        {"2019-01-04"},
        {"2019-01-05"},
    };
    return A_B;
  }

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

  DS filter_beta(Predicate p)
  {
    DS result = *this;
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
  Relation A = Relation::read_from_file("A.txt");
  Relation B = Relation::read_from_file("B.txt");

  cout << A.to_string() << endl;
  cout << B.to_string() << endl;

  DS A_B = DS::make(A, B);

  cout << A_B.to_string() << endl;

  cout << "INITIAL RELATION:\n";
  cout << A_B.to_string() << '\n';

  DS filtered = A_B.filter_beta([](auto owner, auto member, auto row)
                                { return owner[2] == "asthma" || stoi(member[3]) > 50; });

  cout << "FILTERED RELATION:\n";
  cout << filtered.to_string() << '\n';
  // ofstream file("a_b.txt");
  // file << filtered.to_string();

  return 0;
}