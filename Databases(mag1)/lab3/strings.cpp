#include <vector>
#include <string>
#include <sstream>

using namespace std;

using Tuple = vector<string>;

class Strings
{
public:
  static Tuple split(string str, char by)
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

public:
  static string join(vector<string> vec, string delim)
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