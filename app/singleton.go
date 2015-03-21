// An example of a singleton in golang use this for managers like the configuration mangaer.
package singleton

type single struct {
        O interface{};
}

var instantiated *single = nil

func New() *single {
        if instantiated == nil {
                instantiated = new(single);
        }
        return instantiated;
}
