#include </app/file.h>
#include </app/strings.cpp>
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
        record.attributes = Strings::split(line, ',');
        break;
      default:
        record.rows.push_back(Strings::split(line, ','));
        break;
      }
    }
    return record;
  }

public:
  string to_string()
  {
    stringstream stringBuilder;
    stringBuilder << name << "\n";
    stringBuilder << Strings::join(attributes, ",");
    stringBuilder << "\n";
    for (auto row : rows)
    {
      stringBuilder << Strings::join(row, ",");
      stringBuilder << "\n";
    }
    return stringBuilder.str();
  }
};
