string before, after;
void addText(string text) {
    before.insert(end(before), begin(text), end(text));
}
int deleteText(int k) {
    k = min(k, (int)before.size());
    before.resize(before.size() - k);
    return k;
}
string cursorLeft(int k) {
    while(k-- && !before.empty()) {
        after.push_back(before.back());
        before.pop_back();
    }
    return before.substr(before.size() - min((int)before.size(), 10));
}
string cursorRight(int k) {
    while(k-- && !after.empty()) {
        before.push_back(after.back());
        after.pop_back();
    }
    return before.substr(before.size() - min((int)before.size(), 10));
}

int main()
{
    addText("hello");
    addText("world");
    std::cout << "before is " << before << std::endl;
    std::cout << "after is " << after << std::endl;
    std::cout << "moving cursor left 1" << std::endl;
    std::string curr = cursorLeft(1);
    std::cout << curr << std::endl;
    std::cout << "before is " << before << std::endl;
    std::cout << "after is " << after << std::endl;
    std::cout << "moving cursor right 5" << std::endl;
    curr = cursorRight(5);
    std::cout << curr << std::endl;
    
    std::cout << "before is " << before << std::endl;
    std::cout << "after is " << after << std::endl;
    std::cout << "deleting 2 " << std::endl;
    deleteText(2);
    std::cout << "before is " << before << std::endl;
    std::cout << "after is " << after << std::endl;
}