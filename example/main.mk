let map = fn(arr, f) {
	let iter = fn(from, to) {
		if (len(from) == 0) {
			return to;
		} else {
			return iter(rest(from), push(to, f(first(from))));
		}
	}

	iter(arr, [])
};

let a = [1, 2, 3, 4];
let double = fn(x) { x * 2};
map(a, double);
