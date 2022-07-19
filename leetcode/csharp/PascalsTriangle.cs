
namespace PascalsTriangleNamespace
{
    public class PascalsTriangleSolution
    {
        public IList<IList<int>> Generate(int numRows)
        {
            IList<IList<int>> res = new List<IList<int>>();
            if (numRows == 0)
            {
                return res;
            }
            IList<int> first = new List<int>();
            first.Add(1);
            res.Add(first);

            for (int i = 1; i < numRows; i++)
            {

                IList<int> row = new List<int>();
                for (int j = 0; j <= i; j++)
                {
                    if (j == 0 || j == i)
                    {
                        row.Add(1);
                    }
                    else
                    {
                        row.Add(res[i - 1][j - 1] + res[i - 1][j]);
                    }
                }
                res.Add(row);
            }
            return res;
        }
    }
}