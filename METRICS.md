# HELP docker_container_cpu_throttled_time Count of milliseconds when trottling cpu for conrainer is enabled.
# TYPE docker_container_cpu_throttled_time gauge
docker_container_cpu_throttled_time{container_id="0185c775e343",container_name="dev-consul",image="consul"} 0
# HELP docker_container_cpu_usage_ratio Current percent of CPU usage per container.
# TYPE docker_container_cpu_usage_ratio gauge
docker_container_cpu_usage_ratio{container_id="0185c775e343",container_name="dev-consul",image="consul"} 1.583
# HELP docker_container_memory_limit_bytes Current limit bytes of memory usage per container.
# TYPE docker_container_memory_limit_bytes gauge
docker_container_memory_limit_bytes{container_id="0185c775e343",container_name="dev-consul",image="consul"} 6.752854016e+10
# HELP docker_container_memory_usage_bytes Current bytes of memory usage per container.
# TYPE docker_container_memory_usage_bytes gauge
docker_container_memory_usage_bytes{container_id="0185c775e343",container_name="dev-consul",image="consul"} 5.2555776e+07
# HELP docker_container_memory_usage_ratio Current percent of memory usage per container.
# TYPE docker_container_memory_usage_ratio gauge
docker_container_memory_usage_ratio{container_id="0185c775e343",container_name="dev-consul",image="consul"} 0.078
# HELP docker_container_restarts_count Count of container restarts.
# TYPE docker_container_restarts_count gauge
docker_container_restarts_count{container_id="0185c775e343",container_name="dev-consul",image="consul"} 0
# HELP docker_container_state ID of current container state
# TYPE docker_container_state gauge
docker_container_state{container_id="0185c775e343",container_name="dev-consul",image="consul"} 2
# HELP docker_exporter_fetch_metrics_time_ms Time of fetching all metrics from docker daemon(s)
# TYPE docker_exporter_fetch_metrics_time_ms gauge
docker_exporter_fetch_metrics_time_ms 2000






# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 7.7097e-05
go_gc_duration_seconds{quantile="0.25"} 0.000186659
go_gc_duration_seconds{quantile="0.5"} 0.000299242
go_gc_duration_seconds{quantile="0.75"} 0.000894978
go_gc_duration_seconds{quantile="1"} 0.011914862
go_gc_duration_seconds_sum 0.816231487
go_gc_duration_seconds_count 1004
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 77
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 2.54596e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 1.43954604e+09
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 1.535002e+06
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 1.050559e+07
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 544768
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 2.54596e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 4.120576e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 4.661248e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 17428
# HELP go_memstats_heap_released_bytes_total Total number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes_total counter
go_memstats_heap_released_bytes_total 0
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 8.781824e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.5186138768992136e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 33431
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 1.0523018e+07
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 13888
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 80408
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 114688
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.978416e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 2.190558e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 1.703936e+06
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 1.703936e+06
# HELP go_memstats_sys_bytes Number of bytes obtained by system. Sum of all system allocations.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 1.488716e+07