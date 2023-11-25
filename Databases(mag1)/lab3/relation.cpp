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
      cout << "line\n";
      cout << line;
      cout << "line\n";
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

  string to_string()
  {
    stringstream stringBuilder;
    stringBuilder << name << "\n";
    for (auto attribute : attributes)
    {
      stringBuilder << attribute << ",";
    }
    stringBuilder << "\n";
    for (auto row : rows)
    {
      for (auto value : row)
      {
        stringBuilder << value << ",";
      }
      stringBuilder << "\n";
    }
    return stringBuilder.str();
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
};
