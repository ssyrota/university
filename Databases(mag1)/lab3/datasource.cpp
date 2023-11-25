#include </app/relation.cpp>
#include <vector>
#include <string>
#include <sstream>
#include <functional>

using namespace std;

// owner, member, ds row
using Predicate = function<bool(vector<string>, vector<string>, vector<string>)>;

struct DataSource : Relation
{
  Relation owner;
  Relation member;
  bool single_owner;
  bool single_member;
  vector<pair<int, int>> id_pairs;

  static DataSource new_default(Relation a, Relation b)
  {
    DataSource a_b = {
        .owner = a,
        .member = b,
        .single_owner = false,
        .single_member = false,
        .id_pairs = {{1, 1}, {1, 2}, {2, 3}, {1, 3}},
    };
    a_b.name = "a_b";
    a_b.attributes = {"date"};
    a_b.rows = {
        {"2019-01-01"},
        {"2019-01-02"},
        {"2019-01-03"},
        {"2019-01-04"},
        {"2019-01-05"},
    };
    return a_b;
  }

  string to_string()
  {
    stringstream ss;
    ss << name << '\n';
    ss << owner.name << "," << member.name << "," << single_owner << "," << single_member;

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

  DataSource filter_beta(Predicate p)
  {
    DataSource result = {
        .owner = this->owner,
        .member = this->member,
        .single_owner = this->single_owner,
        .single_member = this->single_member,
        .id_pairs = {},
    };
    for (int i = 0; i < id_pairs.size(); i++)
    {
      int owner_id = id_pairs[i].first;
      int member_id = id_pairs[i].second;
      if (p(this->owner.rows[owner_id], this->owner.rows[member_id], this->rows[i]))
      {
        result.rows.push_back(this->rows[i]);
        result.id_pairs.push_back(this->id_pairs[i]);
      }
    }
    return result;
  }

private:
  static void printVector(vector<string> vec)
  {
    for (string str : vec)
    {
      cout << str << " ";
    }
    cout << endl;
  }

public:
  void to_file()
  {
    write_file("./" + name + ".txt", to_string());
  }
};
