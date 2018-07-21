// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package storage

import (
	metrics "github.com/polar9527/go-nebulas/metrics"
)

// Metrics for storage
var (
	// rocksdb metrics
	metricsRocksdbFlushTime = metrics.NewGauge("neb.rocksdb.flushtime")
	metricsRocksdbFlushLen  = metrics.NewGauge("neb.rocksdb.flushlen")

	metricsBlocksdbCacheSize       = metrics.NewGauge("neb.rocksdb.cache.size")       //block_cache->GetUsage()
	metricsBlocksdbCachePinnedSize = metrics.NewGauge("neb.rocksdb.cachepinned.size") //block_cache->GetPinnedUsage()
	metricsBlocksdbTableReaderMem  = metrics.NewGauge("neb.rocksdb.tablereader.mem")  //estimate-table-readers-mem
	metricsBlocksdbAllMemTables    = metrics.NewGauge("neb.rocksdb.alltables.mem")    //cur-size-all-mem-tables
)
