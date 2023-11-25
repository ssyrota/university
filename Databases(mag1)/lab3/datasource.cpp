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
  Relation *owner;
  Relation *member;
  bool single_owner;
  bool single_member;
  vector<pair<int, int>> id_pairs;

  static DataSource new_default(Relation a, Relation b)
  {
    DataSource a_b = {
        .owner = &a,
        .member = &b,
        .single_owner = false,
        .single_member = false,
        .id_pairs = {{1, 1}, {1, 2}, {2, 3}, {1, 3}, {3, 4}},
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

  DataSource filter_beta(Predicate p)
  {
    DataSource result = *this;
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

public:
  void to_file()
  {
    write_file("./" + name + ".txt", to_string());
  }
};
