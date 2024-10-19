package presentation

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n == 1 {
		for i := 0; i < 10; i++ {
			if i == 9 {
				PrintNbr(i)
				PrintStr("\n")
				break
			}
			PrintNbr(i)
			PrintStr(", ")
		}
	}
	if n == 2 {
		for i := '0'; i <= '8'; i++ {
			for j := '1'; j <= '9'; j++ {
				if i < j {
					if i == '8' && j == '9' {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(10)
					} else {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
	if n == 3 {
		PrintComb()
	}
	if n == 4 {
		for i := '0'; i <= '6'; i++ {
			for j := '1'; j <= '7'; j++ {
				for k := '2'; k <= '8'; k++ {
					for l := '3'; l <= '9'; l++ {
						if i < j && j < k && k < l {
							if i == '6' && j == '7' && k == '8' && l == '9' {
								z01.PrintRune(i)
								z01.PrintRune(j)
								z01.PrintRune(k)
								z01.PrintRune(l)
								z01.PrintRune('\n')
							} else {
								z01.PrintRune(i)
								z01.PrintRune(j)
								z01.PrintRune(k)
								z01.PrintRune(l)
								PrintStr(", ")
							}
						}
					}
				}
			}
		}
	}
	if n == 5 {
		for i := '0'; i <= '5'; i++ {
			for j := '1'; j <= '6'; j++ {
				for k := '2'; k <= '7'; k++ {
					for l := '3'; l <= '8'; l++ {
						for m := '4'; m <= '9'; m++ {
							if i < j && j < k && k < l && l < m {
								if i == '5' && j == '6' && k == '7' && l == '8' && m == '9' {
									z01.PrintRune(i)
									z01.PrintRune(j)
									z01.PrintRune(k)
									z01.PrintRune(l)
									z01.PrintRune(m)
									z01.PrintRune('\n')
								} else {
									z01.PrintRune(i)
									z01.PrintRune(j)
									z01.PrintRune(k)
									z01.PrintRune(l)
									z01.PrintRune(m)
									PrintStr(", ")
								}
							}
						}
					}
				}
			}
		}
	}
	if n == 6 {
		for i := '0'; i <= '4'; i++ {
			for j := '1'; j <= '5'; j++ {
				for k := '2'; k <= '6'; k++ {
					for l := '3'; l <= '7'; l++ {
						for m := '4'; m <= '8'; m++ {
							for n := '5'; n <= '9'; n++ {
								if i < j && j < k && k < l && l < m && m < n {
									if i == '4' && j == '5' && k == '6' && l == '7' && m == '8' && n == '9' {
										z01.PrintRune(i)
										z01.PrintRune(j)
										z01.PrintRune(k)
										z01.PrintRune(l)
										z01.PrintRune(m)
										z01.PrintRune(n)
										z01.PrintRune('\n')
									} else {
										z01.PrintRune(i)
										z01.PrintRune(j)
										z01.PrintRune(k)
										z01.PrintRune(l)
										z01.PrintRune(m)
										z01.PrintRune(n)
										PrintStr(", ")
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if n == 7 {
		for i := '0'; i <= '3'; i++ {
			for j := '1'; j <= '4'; j++ {
				for k := '2'; k <= '5'; k++ {
					for l := '3'; l <= '6'; l++ {
						for m := '4'; m <= '7'; m++ {
							for n := '5'; n <= '8'; n++ {
								for o := '6'; o <= '9'; o++ {
									if i < j && j < k && k < l && l < m && m < n && n < o {
										if i == '3' && j == '4' && k == '5' && l == '6' && m == '7' && n == '8' && o == '9' {
											z01.PrintRune(i)
											z01.PrintRune(j)
											z01.PrintRune(k)
											z01.PrintRune(l)
											z01.PrintRune(m)
											z01.PrintRune(n)
											z01.PrintRune(o)
											z01.PrintRune('\n')
										} else {
											z01.PrintRune(i)
											z01.PrintRune(j)
											z01.PrintRune(k)
											z01.PrintRune(l)
											z01.PrintRune(m)
											z01.PrintRune(n)
											z01.PrintRune(o)
											PrintStr(", ")
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if n == 8 {
		for i := '0'; i <= '2'; i++ {
			for j := '1'; j <= '3'; j++ {
				for k := '2'; k <= '4'; k++ {
					for l := '3'; l <= '5'; l++ {
						for m := '4'; m <= '6'; m++ {
							for n := '5'; n <= '7'; n++ {
								for o := '6'; o <= '8'; o++ {
									for p := '7'; p <= '9'; p++ {
										if i < j && j < k && k < l && l < m && m < n && n < o && o < p {
											if i == '2' && j == '3' && k == '4' && l == '5' && m == '6' && n == '7' && o == '8' && p == '9' {
												z01.PrintRune(i)
												z01.PrintRune(j)
												z01.PrintRune(k)
												z01.PrintRune(l)
												z01.PrintRune(m)
												z01.PrintRune(n)
												z01.PrintRune(o)
												z01.PrintRune(p)
												z01.PrintRune('\n')
											} else {
												z01.PrintRune(i)
												z01.PrintRune(j)
												z01.PrintRune(k)
												z01.PrintRune(l)
												z01.PrintRune(m)
												z01.PrintRune(n)
												z01.PrintRune(o)
												z01.PrintRune(p)
												PrintStr(", ")
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if n == 9 {
		for i := '0'; i <= '1'; i++ {
			for j := '1'; j <= '2'; j++ {
				for k := '2'; k <= '3'; k++ {
					for l := '3'; l <= '4'; l++ {
						for m := '4'; m <= '5'; m++ {
							for n := '5'; n <= '6'; n++ {
								for o := '6'; o <= '7'; o++ {
									for p := '7'; p <= '8'; p++ {
										for q := '8'; q <= '9'; q++ {
											if i < j && j < k && k < l && l < m && m < n && n < o && o < p && p < q {
												if i == '1' && j == '2' && k == '3' && l == '4' && m == '5' && n == '6' && o == '7' && p == '8' && q == '9' {
													z01.PrintRune(i)
													z01.PrintRune(j)
													z01.PrintRune(k)
													z01.PrintRune(l)
													z01.PrintRune(m)
													z01.PrintRune(n)
													z01.PrintRune(o)
													z01.PrintRune(p)
													z01.PrintRune(q)
													z01.PrintRune('\n')
												} else {
													z01.PrintRune(i)
													z01.PrintRune(j)
													z01.PrintRune(k)
													z01.PrintRune(l)
													z01.PrintRune(m)
													z01.PrintRune(n)
													z01.PrintRune(o)
													z01.PrintRune(p)
													z01.PrintRune(q)
													PrintStr(", ")
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
