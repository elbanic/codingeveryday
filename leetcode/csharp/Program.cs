using HelloWorldNamespace;
using PascalsTriangleNamespace;

namespace CodingEveryDay
{
    class Program {         
        static void Main(string[] args)
        {
            HelloWorld h = new HelloWorld();
            h.print();

            PascalsTriangleSolution pt = new PascalsTriangleSolution();
            IList<IList<int>> res = pt.Generate(5);

        }
    }
}