package category

type Service interface {
	// Mengambil semua data kategori
	GetAllCategory() (*[]Category, error)

	// Mengambil data kategori berdasarkan kode kategori
	FindCategoryById(id string) (*Category, error)

	FindCategoryByCode(code string) (*Category, error)

	// Menambahan kategori baru
	InsertCategory(category *CategorySpec, creator string) error

	// Memperbarui produk kategori
	UpdateCategory(id string, category *CategorySpec, modifier string) error

	// Menghapus produk kategori
	DeleteCategory(id string, deleter string) error
}

type Repository interface {
	// Mengambil semua data kategori
	GetAllCategory() (*[]Category, error)

	// Mengambil data kategoru bedasarkan kode kategori
	FindCategoryById(id string) (*Category, error)

	FindCategoryByCode(code string) (*Category, error)

	// Menambahkan kategori baru
	InsertCategory(category Category) error

	// Memperbarui kategori
	UpdateCategory(id string, category Category) error

	// Menghapus produk kategori
	DeleteCategory(id string, adminId string) error
}
