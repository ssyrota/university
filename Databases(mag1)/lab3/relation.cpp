#include </app/file.h>
#include <vector>
#include <fstream>
#include <string>
#include <sstream>

using namespace std;

using Tuple = vector<string>;

struct Relation
{
  string name;
  Tuple attributes;
  vector<Tuple> rows;

  static Relation make_from_file(string file_name)
  {
    string file = read_file(file_name);
    Relation record;
    istringstream linesStream(file);
    string line;
    int lineNumber = 0;
    while (getline(linesStream, line, '\n'))
    {
      lineNumber++;
      switch (lineNumber)
      {
      case 1:
        record.name = line;
        break;
      case 2:
        record.attributes = Relation::splitString(line, ',');
        break;
      default:
        record.rows.push_back(Relation::splitString(line, ','));
        break;
      }
    }
    return record;
  }

private:
  string to_string()
  {
    stringstream stringBuilder;
    stringBuilder << name << "\n";
    stringBuilder << Relation::join_string(attributes, ",");
    stringBuilder << "\n";
    for (auto row : rows)
    {
      stringBuilder << Relation::join_string(row, ",");
      stringBuilder << "\n";
    }
    return stringBuilder.str();
  }

public:
  void to_file()
  {
    write_file("./" + name + ".txt", to_string());
  }

  static Tuple splitString(string str, char by)
  {
    Tuple result;
    istringstream iss(str);
    string token;
    while (getline(iss, token, by))
    {
      result.push_back(token);
    }
    return result;
  }

  static string join_string(vector<string> vec, string delim)
  {
    string result;
    for (size_t i = 0; i < vec.size(); ++i)
    {
      result += vec[i];
      if (i < vec.size() - 1)
      {
        result += delim;
      }
    }
    return result;
  }
};
