# Bracket Balance

## Complexity
### Time
Time complexity dari solusi ini adalah O(n), di mana n adalah panjang string input. Ini karena kita hanya meng-iterasi melalui string sekali dan setiap operasi push dan pop ke dan dari stack adalah O(1).

### Space
Worst case dari space complexity-nya juga O(n), itu terjadi ketika semua karakter dalam string adalah tanda kurung pembuka sehingga semuanya masuk ke dalam stack (map).