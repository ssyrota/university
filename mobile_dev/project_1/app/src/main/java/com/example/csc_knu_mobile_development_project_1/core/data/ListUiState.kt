package com.example.csc_knu_mobile_development_project_1.core.data

data class ListUiState(
	var list: MutableList<Double> = mutableListOf(),
)

/*TODO: add cycles count*/
class SortedList(private val input: List<Double>) {
	fun value(): List<Double> {
		return this.input.sorted()
	}

	fun sortStats(): SortOpsCount {
		val stats = HashMap<SortType, Int>()
		stats.set(SortType.INSERTION, this.insertionSort())
		stats.set(SortType.SELECTION, this.selectionSort())
		stats.set(SortType.QUICK, this.quickSort())
		stats.set(SortType.PANCAKE, this.pancakeSort())
		stats.set(SortType.MERGE, this.mergeSort().second)
		return stats
	}

	private fun selectionSort(): Int {
		val items = input.toMutableList()
		var n = items.count()
		var temp: Double
		var ops = 0
		for (i in 0 until n) {
			ops++
			var indexOfMin = i
			for (j in n - 1 downTo i) {
				ops++
				if (items[j] < items[indexOfMin])
					indexOfMin = j
			}
			if (i != indexOfMin) {
				temp = items[i]
				items[i] = items[indexOfMin]
				items[indexOfMin] = temp
			}
		}
		return ops
	}

	private fun insertionSort(): Int {
		val items = input.toMutableList()
		if (items.count() < 2) {
			return 2
		}

		var ops = 0
		for (count in 1 until items.count()) {
			ops++
			val item = items[count]
			var i = count
			while (i > 0 && item < items[i - 1]) {
				ops++
				items[i] = items[i - 1]
				i -= 1
			}
			items[i] = item
		}
		return ops
	}

	private fun quickSort(
		items: List<Double> = input.toList()
	): Int {
		if (items.count() < 2) {
			return 2
		}
		val pivot = items[items.count() / 2]
		val equal = items.filter { it == pivot }
		val less = items.filter { it < pivot }
		val greater = items.filter { it > pivot }
		return quickSort(
			less
		) + items.size + quickSort(greater)
	}

	private fun mergeSort(items: List<Double> = input.toList()): Pair<List<Double>, Int> {
		if (items.size <= 1) {
			return Pair(items, 1)
		}
		val middle = items.size / 2
		var left = items.subList(0, middle);
		var right = items.subList(middle, items.size);
		return merge(mergeSort(left), mergeSort(right))
	}

	private fun merge(
		left: Pair<List<Double>, Int>,
		right: Pair<List<Double>, Int>
	): Pair<List<Double>, Int> {
		var ops = left.second + right.second
		var indexLeft = 0
		var indexRight = 0
		var newList = mutableListOf<Double>()
		while (indexLeft < left.first.count() && indexRight < right.first.count()) {
			ops++
			if (left.first[indexLeft] <= right.first[indexRight]) {
				newList.add(left.first[indexLeft])
				indexLeft++
			} else {
				newList.add(right.first[indexRight])
				indexRight++
			}
		}
		while (indexLeft < left.first.size) {
			ops++
			newList.add(left.first[indexLeft])
			indexLeft++
		}
		while (indexRight < right.first.size) {
			ops++
			newList.add(right.first[indexRight])
			indexRight++
		}
		return Pair(newList, ops)
	}

	private fun pancakeSort(items: MutableList<Double> = input.toMutableList()): Int {
		var ops = 0
		for (n in items.count() downTo 2) {
			ops++
			val (maxI, indexOfMaxOps) = this.indexOfMax(items, n)
			ops += indexOfMaxOps
			if (maxI != n - 1) {
				if (maxI > 0) {
					ops += pancakeFlipToStart(items, maxI)
				}
				ops += pancakeFlipToStart(items, n - 1)
			}
		}
		return ops
	}

	private fun indexOfMax(items: List<Double>, n: Int): Pair<Int, Int> {
		var ops = 0
		var index = 0
		for (i in 1 until n) {
			ops++
			if (items[i] > items[index]) index = i
		}
		return Pair(index, ops)
	}

	private fun pancakeFlipToStart(items: MutableList<Double>, index: Int): Int {
		var ops = 0
		var i = index
		var j = 0
		while (j < i) {
			ops++
			val temp = items[j]
			items[j] = items[i]
			items[i] = temp
			j++
			i--
		}
		return ops
	}
}

enum class SortType(name: String) {
	PANCAKE("Pancake"),
	SELECTION("Selection"),
	INSERTION("Insertion"),
	QUICK("Quick"),
	MERGE("Merge")
}
typealias SortOpsCount = Map<SortType, Int>